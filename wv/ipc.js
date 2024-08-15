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
    const MT_READY = 1;
    const MT_EVENT_GO_EMIT = MT_READY + 1;
    const MT_EVENT_JS_EMIT = MT_EVENT_GO_EMIT + 1;
    const MT_EVENT_GO_EMIT_CALLBACK = MT_EVENT_JS_EMIT + 1;
    const MT_EVENT_JS_EMIT_CALLBACK = MT_EVENT_GO_EMIT_CALLBACK + 1;
    const MT_DRAG_MOVE = MT_EVENT_JS_EMIT_CALLBACK + 1;
    const MT_DRAG_DOWN = MT_DRAG_MOVE + 1;
    const MT_DRAG_UP = MT_DRAG_DOWN + 1;
    const MT_DRAG_DBLCLICK = MT_DRAG_UP + 1;
    const MT_DRAG_RESIZE = MT_DRAG_DBLCLICK + 1;
    const MT_DRAG_BORDER_WMSZ = MT_DRAG_RESIZE + 1;

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

        //drag
        #drag;

        #env = {};

        /**
         * js process message
         * @param {string} message json
         * @public
         */
        processMessage(message) {
            throw new Error("Unsupported Platform");
        };

        /**
         * energy set env arguments
         * @param {string} key
         * @param {string} value
         * @public
         */
        setEnv(key, value) {
            this.#env[key] = value
        }

        /**
         * energy options set env
         * @param {JSON} options
         * @public
         */
        setOptionsEnv(options) {
            for (let key in options) {
                this.setEnv(key, options[key]);
            }
        }

        /**
         * energy get env arguments
         * @param {string} key
         * @public
         */
        getEnv(key) {
            return this.#env[key]
        }

        /**
         * Creates an instance of Energy.
         * @memberof Energy
         */
        constructor() {
            this.#eventListeners = new Map();
            this.#emitCallbacks = new Map();
            this.#executionID = 0;
            // process message
            if (this.#deepTest(["chrome", "webview", "postMessage"])) {
                // webview2
                let webview = window.chrome.webview;
                // render process send message => go
                this.processMessage = (message) => webview.postMessage(message);
                // render process receive browser process string message
                webview.addEventListener("message", event => {
                    window.energy.__executeEvent(event.data);
                });
                // render process receive browser process buffer message
                //webview.addEventListener("sharedbufferreceived", event => {
                //let buffer = event.getBuffer();
                //let bufferData = new TextDecoder().decode(new Uint8Array(buffer));
                // console.log("buffer:", bufferData);
                //});
            } else if (this.#deepTest(["webkit", "messageHandlers", "external", "postMessage"])) {
                // webkit
                // render process send message => go
                this.processMessage = (message) => window.webkit.messageHandlers.external.postMessage(message);
            } else {
                throw new Error("Unsupported Platform");
            }
            this.#drag = new Drag();
        }

        drag() {
            return this.#drag;
        }

        #deepTest(s) {
            let obj = window[s.shift()];
            while (obj && s.length) obj = obj[s.shift()];
            return obj;
        };

        /**
         * @param {object} message
         */
        #notifyListeners(message) {
            switch (message.t) {
                case MT_EVENT_GO_EMIT:
                    this.#handlerGOEMIT(message);
                    break
                case MT_EVENT_JS_EMIT_CALLBACK:
                    this.#handlerJSEMITCallback(message);
                    break
            }
        };

        /**
         * @param {object} message
         */
        #handlerJSEMITCallback(message) {
            let id = message.i;                         // executionID
            let callback = this.#emitCallbacks.get(id); // get ipc.emit callback function
            if (callback) {
                this.#emitCallbacks.delete(id); // remove ipc.emit callback function by executionID
                let args = message.d;           // arguments
                if (!Array.isArray(args)) {
                    args = [args];
                }
                callback.apply(null, args);
            }
        }

        /**
         * @param {object} message
         */
        #handlerGOEMIT(message) {
            let id = message.i;   // executionID
            let name = message.n; // name
            let callback = this.#eventListeners.get(name);
            if (callback) {
                let args = message.d; // arguments
                if (!Array.isArray(args)) {
                    args = [args];
                }
                let result = callback.apply(null, args);
                // not 0 go has callback function
                if (id !== 0) {
                    const payload = {
                        t: MT_EVENT_GO_EMIT_CALLBACK,                  // MessageType
                        n: name,                                 // name
                        d: [].slice.apply([result]),     // data
                        i: id,                                   // executionID
                    };
                    this.processMessage(JSON.stringify(payload));
                }
            }
        }

        /**
         * @param {string} name
         * @param {function} callback
         * @private
         */
        __setEventListener(name, callback) {
            this.#eventListeners.set(name, callback);
        }

        /**
         * @param {string} name
         * @private
         */
        __removeEventListener(name) {
            this.#eventListeners.delete(name);
        }

        /**
         * @param {number} executionID
         * @param {function} callback
         * @private
         */
        __setJSEmitCallback(executionID, callback) {
            this.#emitCallbacks.set(executionID, callback);
        }

        /**
         * @param {string} messageData
         * @private
         */
        __executeEvent(messageData) {
            try {
                this.#notifyListeners(JSON.parse(messageData));
            } catch (e) {
                throw new Error(e + ' ' + messageData);
            }
        };

        /**
         * return the ID of the next IPC message executed in JavaScript
         * @returns {number} messageId
         * @private
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
         */
        removeOn(name) {
            window.energy.__removeEventListener(name);
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
                t: MT_EVENT_JS_EMIT,           // MessageType
                n: name,                 // name
                d: [].slice.apply(data), // data
                i: executionID,          // executionID
            };
            // call js event

            // call go event
            energy.processMessage(JSON.stringify(payload));
        }
    }

    class Drag {
        #shouldDrag = false;
        #cssDragProperty = "-webkit-app-region";
        #cssDragValue = "drag";

        constructor() {
        }

        #test(e) {
            let v = window.getComputedStyle(e.target)[this.#cssDragProperty];
            if (v) {
                v = v.trim();
                if (v !== this.#cssDragValue) {
                    return false;
                }
                // return e.buttons === 1;
                return e.detail === 1 || e.detail === 2;
            }
            return false;
        }

        setup() {
            console.log("setup");
            let that = this;
            function dragMessage(t, n, d) {
                const payload = {t: t, n: n, d: d};
                energy.processMessage(JSON.stringify(payload));
            }

            let idcCursor = null;
            let frameWidth = energy.getEnv("frameWidth") || 4;
            let frameHeight = energy.getEnv("frameHeight") || 4;
            let frameCorner = energy.getEnv("frameCorner") || 8;
            let disableResize = energy.getEnv("disableResize") || false;
            let disableWebkitAppRegionDClk = energy.getEnv("disableWebkitAppRegionDClk") || false;
            let isWindows = energy.getEnv("os") === "windows";
            let frameless = energy.getEnv("frameless") || false;

            function setCursor(cursor, ht) {
                if (idcCursor !== cursor) {
                    document.documentElement.style.cursor = cursor || 'auto';
                    idcCursor = cursor;
                    // dragMessage(MT_DRAG_BORDER_WMSZ, 'wmsz', ht);
                }
            }

            function mouseDragResize(e) {
                let leftBorder = e.clientX < frameWidth;
                let topBorder = e.clientY < frameHeight;
                let rightBorder = window.outerWidth - e.clientX < frameWidth;
                let bottomBorder = window.outerHeight - e.clientY < frameHeight;
                let leftCorner = e.clientX < frameWidth + frameCorner;
                let topCorner = e.clientY < frameHeight + frameCorner;
                let rightCorner = window.outerWidth - e.clientX < frameWidth + frameCorner;
                let bottomCorner = window.outerHeight - e.clientY < frameHeight + frameCorner;
                if (!leftBorder && !topBorder && !rightBorder && !bottomBorder && idcCursor !== void 0) {
                    setCursor();
                } else if (rightCorner && bottomCorner) {
                    setCursor("se-resize", 17);
                } else if (leftCorner && bottomCorner) {
                    setCursor("sw-resize", 16);
                } else if (leftCorner && topCorner) {
                    setCursor("nw-resize", 13);
                } else if (topCorner && rightCorner) {
                    setCursor("ne-resize", 14);
                } else if (leftBorder) {
                    setCursor("w-resize", 10);
                } else if (topBorder) {
                    setCursor("n-resize", 12);
                } else if (bottomBorder) {
                    setCursor("s-resize", 15);
                } else if (rightBorder) {
                    setCursor("e-resize", 11);
                }
            }

            function mouseMove(e) {
                if (that.#shouldDrag) {
                    if (isWindows) {
                        that.#shouldDrag = false;
                    }
                    dragMessage(MT_DRAG_MOVE, 'move', {x: e.screenX, y: e.screenY});
                } else if (!disableResize && isWindows && frameless) {
                    mouseDragResize(e)
                }
            }

            function mouseUp(e) {
                that.#shouldDrag = false;
                if (that.#test(e)) {
                    e.preventDefault();
                    dragMessage(MT_DRAG_UP, 'up', null);
                }
            }

            function mouseDown(e) {
                if (idcCursor) {
                    e.preventDefault();
                    dragMessage(MT_DRAG_RESIZE, 'resize', idcCursor);
                } else if (!(e.offsetX > e.target.clientWidth || e.offsetY > e.target.clientHeight) && that.#test(e)) {
                    e.preventDefault();
                    that.#shouldDrag = true;
                    dragMessage(MT_DRAG_DOWN, 'down', {x: e.screenX, y: e.screenY});
                } else {
                    that.#shouldDrag = false;
                }
            }

            function dblClick(e) {
                if (that.#test(e) && !disableWebkitAppRegionDClk) {
                    e.preventDefault();
                    dragMessage(MT_DRAG_DBLCLICK, 'dblclk', null);
                }
            }

            window.addEventListener("mousemove", mouseMove);
            window.addEventListener("mousedown", mouseDown);
            window.addEventListener("mouseup", mouseUp);
            window.addEventListener("dblclick", dblClick);
        }
    }

    window.energy = new Energy();
    window.ipc = new IPC();
    //window.energy.processMessage(JSON.stringify({t: MT_READY, n: 'ready'}));
    console.log("ipc.js");
})();
