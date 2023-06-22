type Callback = (...args: any) => void;
const events = new Map<string, Callback[]>();

export function useEventBus() {
  const on = (event: string, cb: Callback) => {
    let list = events.get(event);
    if (list) {
      list.push(cb);
    } else {
      list = [cb];
      events.set(event, list);
    }
  };

  const emit = (event: string, ...data: any) => {
    const list = events.get(event);
    if (!list) return;
    for (const cb of list) {
      cb(data);
    }
  };

  const clear = (event: string) => {
    events.delete(event);
  };

  return {
    on,
    emit,
    clear,
  };
}
