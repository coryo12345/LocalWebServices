type Callback = (...args: any) => void;
const eventMap = new Map<string, Callback[]>();

export function useEventBus() {
  const on = (event: string, cb: Callback) => {
    let list = eventMap.get(event);
    if (list) {
      list.push(cb);
    } else {
      list = [cb];
      eventMap.set(event, list);
    }
  };

  const emit = (event: string, ...data: any) => {
    const list = eventMap.get(event);
    if (!list) return;
    for (const cb of list) {
      cb(...data);
    }
  };

  const clear = (event: string) => {
    eventMap.delete(event);
  };

  return {
    on,
    emit,
    clear,
  };
}

export const events = {
  SNACKBAR_SUCCESS: "snackbar-event-success",
  SNACKBAR_ERROR: "snackbar-event-error",
  SNACKBAR_INFO: "snackbar-event-info",
  SNACKBAR_WARNING: "snackbar-event-warning",
};
