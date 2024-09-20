import { useSignal } from "@preact/signals";
import { useEventBus, events } from "../../utils/bus";
import { useEffect } from "preact/hooks";

enum MessageType {
  SUCCESS,
  ERROR,
  WARNING,
  INFO,
}

const DISPLAY_TIMEOUT_MS = 5000;

export function Snackbar() {
  const { on } = useEventBus();

  useEffect(() => {
    on(events.SNACKBAR_SUCCESS, (msg: string) =>
      addMsg(MessageType.SUCCESS, msg)
    );
    on(events.SNACKBAR_ERROR, (msg: string) => addMsg(MessageType.ERROR, msg));
    on(events.SNACKBAR_INFO, (msg: string) => addMsg(MessageType.INFO, msg));
    on(events.SNACKBAR_WARNING, (msg: string) =>
      addMsg(MessageType.WARNING, msg)
    );
  }, []);

  const messages = useSignal<{ type: MessageType; msg: string }[]>([]);
  function addMsg(type: MessageType, msg: string) {
    messages.value = [...messages.value, { type, msg }];
    const idx = messages.value.length - 1;
    setTimeout(() => {
      messages.value = messages.value.filter((_, i) => i !== idx);
    }, DISPLAY_TIMEOUT_MS);
  }

  function typeColor(type: MessageType): string {
    switch (type) {
      case MessageType.SUCCESS:
        return "bg-lime-500 text-white";
      case MessageType.ERROR:
        return "bg-red-500 text-white";
      case MessageType.INFO:
        return "bg-purple-950 text-white";
      case MessageType.WARNING:
        return "bg-orange-500 text-white";
      default:
        return "";
    }
  }

  return (
    <div class="fixed bottom-0 left-0 text-center w-full">
      {messages.value.map((item) => (
        <div
          class={`mx-auto mb-2 w-fit p-2 rounded border border-gray-400 ${typeColor(
            item.type
          )}`}
        >
          <p>{item.msg}</p>
        </div>
      ))}
    </div>
  );
}
