import { useComputed, useSignal } from "@preact/signals";
import { useContext, useEffect } from "preact/hooks";
import { events, useEventBus } from "../utils/bus";
import { Dialog } from "./base/Dialog";
import { AppState } from "../state";

export const UPDATE_PROPERTY_EVENT = "update-property-dialog-event";

export function UpdatePropertyDialog() {
  const { on, emit } = useEventBus();
  const state = useContext(AppState);

  const dialogValue = useSignal(false);

  const valueText = useSignal("");
  const keys = useSignal<string[]>([]);

  useEffect(() => {
    on(UPDATE_PROPERTY_EVENT, (_keys: string[]) => {
      dialogValue.value = true;
      valueText.value = "";
      keys.value = _keys;
    });
  }, []);

  const submitDisabled = useComputed(() => {
    return !valueText.value.length;
  });

  async function submit() {
    const res = await state.setProperties(keys.value, valueText.value);
    if (res === null) {
      dialogValue.value = false;
      emit(
        events.SNACKBAR_ERROR,
        "Something went wrong, Unable to update properties"
      );
    } else {
      dialogValue.value = false;
      emit(events.SNACKBAR_SUCCESS, "Successfully updated properties");
      await state.fetchProperties();
    }
  }

  return (
    <Dialog
      value={dialogValue.value}
      hideActivator={true}
      title="Update Properties"
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
