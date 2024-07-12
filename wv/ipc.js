//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------


// render process send process message
(function () {
    class Browser {
        #windowId = 0;
        #frameId = 0;

        getWindowId() {
            return this.#windowId;
        }

        getFrameId() {
            return this.#frameId;
        }
    }

    // Energy
    class Energy {
        // js ipc.on event listener
        // @key {string} event name
        // @value {Listener} listener object
        #eventListeners;

        // js ipc.emit callbacks
        // @key {number} executionID
        // @value {function} callback
        #emitCallbacks;

        // js ipc.emit callback executionID, global accumulation
        #executionID;

        /**
         * Creates an instance of Energy.
         * @memberof Energy
         */
        constructor() {
            this.#eventListeners = new Map();
            this.#emitCallbacks = new Map();
            this.#executionID = 0;
        }

        /**
         * @param {object} message
         * @return If there is a return value
         */
        #notifyListeners(message) {
            let id = message.i;
            let name = message.n;
            let callback;
            if (!name && id !== 0) {
                callback = this.#emitCallbacks.get(id);
                if (callback) {
                    this.#emitCallbacks.delete(id);
                }
            } else {
                callback = this.#eventListeners.get(name);
            }
            if (callback) {
                let args = message.d;
                if (!Array.isArray(args)) {
                    args = [args]
                }
                let result = callback.apply(null, args);
                return {
                    id: id,
                    result: result,
                }
            }
        };

        /**
         * @param {string} name
         * @param {function} callback
         */
        __setEventListener(name, callback) {
            this.#eventListeners.set(name, callback);
        }

        /**
         * @param {number} executionID
         * @param {function} callback
         */
        __setJSEmitCallback(executionID, callback) {
            this.#emitCallbacks.set(executionID, callback);
        }

        /**
         * @param {string} messageData
         * @return If there is a return value
         */
        __executeEvent(messageData) {
            try {
                const result = this.#notifyListeners(JSON.parse(messageData));
                if (result) {
                    return result
                }
            } catch (e) {
                throw new Error(e + ' ' + messageData);
            }
        };

        /**
         * return the ID of the next IPC message executed in JavaScript
         * @returns {number} messageId
         */
        __nextExecutionID() {
            this.#executionID++;
            return this.#executionID;
        };
    }

    // IPC
    class IPC {
        /**
         * @param {string} name
         * @param {function} callback
         */
        on(name, callback) {
            if (name && typeof callback === 'function') {
                // __energyEventListeners[name] = __energyEventListeners[name] || [];
                // __energyEventListeners[name].push(thisListener);
                window.energy.__setEventListener(name, callback);
            }
        }

        /**
         * @param {string} name
         * @param {argument} args
         */
        emit(name, ...args) {
            if (!name) {
                throw new Error('ipc.emit call event name is null');
            } else if (args.length > 2) {
                throw new Error('Invalid ipc.emit call arguments');
            }
            let data = [];
            let callback = null;
            let executionID = 0;
            if (args.length === 1) {
                let arg0 = args[0];
                if (Array.isArray(arg0)) {
                    data = arg0;
                } else if (typeof arg0 === 'function') {
                    callback = arg0;
                } else {
                    throw new Error('Invalid ipc.emit call parameter');
                }
            } else if (args.length === 2) {
                let argumentList = args[0]; // array
                let callbackFunc = args[1]; // function
                if (Array.isArray(argumentList) && typeof callbackFunc === 'function') {
                    data = argumentList;
                    callback = callbackFunc;
                } else {
                    throw new Error('Invalid ipc.emit call arguments');
                }
            }
            if (callback !== null) {
                executionID = window.energy.__nextExecutionID();
                window.energy.__setJSEmitCallback(executionID, callback)
            }
            const payload = {
                n: name, // name
                d: [].slice.apply(data), // data
                i: executionID, // executionID
            };
            // call js event

            // call go event
            ProcessMessage(JSON.stringify(payload));
        }
    }

    window.energy = new Energy();
    window.ipc = new IPC();

    let deepTest = function (s) {
        let obj = window[s.shift()];
        while (obj && s.length) obj = obj[s.shift()];
        return obj;
    };
    if (deepTest(["chrome", "webview", "postMessage"])) {
        // webview2
        let webview = window.chrome.webview;
        // render process send message => go
        window.ProcessMessage = (message) => webview.postMessage(message);
        // render process receive browser process string message
        webview.addEventListener("message", event => {
            const result = window.energy.__executeEvent(event.data);
            if (result) {
                console.log("ipc.on-execute result:", result)
            }
        });
        // render process receive browser process buffer message
        webview.addEventListener("sharedbufferreceived", event => {
            let buffer = event.getBuffer();
            let bufferData = new TextDecoder().decode(new Uint8Array(buffer));
            console.log("buffer:", bufferData);
        });
    } else if (deepTest(["webkit", "messageHandlers", "external", "postMessage"])) {
        // webkit
        // render process send message => go
        window.ProcessMessage = (message) => window.webkit.messageHandlers.external.postMessage(message);
    } else {
        console.error("Unsupported Platform");
    }
})();
