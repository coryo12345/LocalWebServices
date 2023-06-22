import { useComputed, useSignal } from "@preact/signals";
import { useEffect } from "preact/hooks";
import { useEventBus } from "../utils/bus";
import { Dialog } from "./base/Dialog";

export const ADD_PROPERTY_EVENT = "add-property-dialog-event";

export function AddPropertyDialog() {
  const { on } = useEventBus();

  const dialogValue = useSignal(false);

  const keyText = useSignal("");
  const valueText = useSignal("");

  useEffect(() => {
    on(ADD_PROPERTY_EVENT, () => {
      dialogValue.value = true;
      keyText.value = "";
      valueText.value = "";
    });
  }, []);

  const submitDisabled = useComputed(() => {
    return !keyText.value.length || !valueText.value.length;
  });

  function submit() {
    // TODO add this property
    console.log(keyText.value, valueText.value);
  }

  return (
    <Dialog
      value={dialogValue.value}
      hideActivator={true}
      title="Add Property"
      onValueChange={(val) => (dialogValue.value = val)}
      actionButtons={
        <button
          disabled={submitDisabled.value}
          class={`py-1 px-2 mr-2 rounded border ${
            submitDisabled.value
              ? "bg-gray-100 text-gray-500 cursor-not-allowed"
              : "bg-blue-700 text-white"
          }`}
          onClick={submit}
        >
          Submit
        </button>
      }
    >
      <div class="flex flex-col">
        <input
          value={keyText.value}
          type="text"
          placeholder="Key"
          id="add-property-key"
          class="border border-black rounded px-2 py-1 mb-2"
          onInput={(e) => (keyText.value = e.currentTarget.value)}
        />
        <input
          value={valueText.value}
          type="text"
          placeholder="Value"
          id="add-property-value"
          class="border border-black rounded px-2 py-1"
          onInput={(e) => (valueText.value = e.currentTarget.value)}
        />
      </div>
    </Dialog>
  );
}
