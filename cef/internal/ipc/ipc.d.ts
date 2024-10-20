/**
 * IPC event driven
 * Listening to/triggering events in JS to achieve interaction with Go
 */
interface IPC {
    /**
     * ipc.emit
     * @param name Event name monitored in Go
     * @param argument The parameter list passed only supports basic data types and composite structures. When receiving in Go, attention should be paid to the parameter types
     * @param callback After the event is triggered, the callback function receives the synchronous return result from Go
     */
    emit(name: string, argument?: any[], callback?: (result: any) => void): void;

    /**
     * ipc.emit
     * @param name Event name monitored in Go
     * @param callback After the event is triggered, the callback function receives the synchronous return result from Go
     */
    emit(name: string, callback?: (result: any) => void): void;

    /**
     * ipc.emit
     */
    emit(rule: EmitRule);

    /**
     * ipc.on
     * @param name Event name monitored in JavaScript
     * @param callback The function called when triggering the JS listening event name in Go, which has parameter reception and return value
     * @param options Listening options, configure listening event behavior
     */
    on(name: string, callback: (...arguments: any[]) => any, options?: Options);
}

/**
 * ipc.emit Default: Mode Trigger asynchronous
 */
declare const EmitModeAsync = 0;

/**
 * ipc.emit Mode Trigger synchronization
 */
declare const EmitModeWait = 0;

/**
 * ipc.emit Default: Target Trigger main process
 */
declare const EmitTargetMain = 0

/**
 * ipc.emit Target Trigger current render process
 */
declare const EmitTargetCurrent = 1

/**
 * ipc.emit Target Trigger other render process
 */
declare const EmitTargetOther = 2

interface EmitRule {
    name: string;
    arguments?: any[];
    callback?: (...result: any[]) => any;
    mode?: number; // EmitModeAsync | EmitModeWait
    target?: number; // EmitTargetMain | EmitTargetCurrent | EmitTargetOther
    waitTime?: number; // default 5000 millisecond
}

/**
 * ipc.on options: mode
 * Default listening mode, indicating that the return value is returned synchronously
 */
declare const MSync = 0;

/**
 * ipc.on options: mode
 * Listening mode, indicating that the return value is returned asynchronously
 * The Complete argument object is automatically added to the last argument in the callback function argument list
 */
declare const MAsync = 1;

/**
 * ipc.on options
 */
interface Options {
    mode: number;
}

/**
 * ipc.on
 * When the listening option mode is MAsync, the Complete object will be automatically added to the last parameter position of the callback function
 */
interface Complete {
    /**
     * Asynchronous return callback function, only valid for the first call
     * @param results result list
     */
    callback: (...results: any[]) => void;
    /**
     * Callback function message ID
     */
    id: number;
}

/**
 * ipc
 */
declare var ipc: IPC;

export {ipc, MSync, MAsync, type Complete, type Options}