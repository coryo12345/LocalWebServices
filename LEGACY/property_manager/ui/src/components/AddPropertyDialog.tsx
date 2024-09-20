import { useComputed, useSignal } from "@preact/signals";
import { useContext, useEffect } from "preact/hooks";
import { events, useEventBus } from "../utils/bus";
import { Dialog } from "./base/Dialog";
import { AppState } from "../state";

export const ADD_PROPERTY_EVENT = "add-property-dialog-event";

export function AddPropertyDialog() {
  const { on, emit } = useEventBus();
  const state = useContext(AppState);

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

  async function submit() {
    const res = await state.addProperty(keyText.value, valueText.value);
    if (res === null) {
      dialogValue.value = false;
      emit(
        events.SNACKBAR_ERROR,
        "Something went wrong, Unable to add new property"
      );
    } else {
      dialogValue.value = false;
      emit(events.SNACKBAR_SUCCESS, "Successfully added new property");
      await state.fetchProperties();
    }
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
