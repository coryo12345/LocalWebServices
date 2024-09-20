import { useComputed, useSignal } from "@preact/signals";
import { JSX, PropsWithChildren, useEffect } from "preact/compat";

interface Props {
  hideActivator?: boolean;
  value?: boolean;
  onValueChange?: (v: boolean) => void;
  actionButtons?: JSX.Element;
  title?: string;
  closeOnClick?: boolean;
}

export function Dialog(props: PropsWithChildren<Props>) {
  const dialogOpen = useSignal(!!props.value);
  const closeOnClick = useComputed(() => !(props.closeOnClick === false));

  useEffect(() => {
    dialogOpen.value = !!props.value;
  }, [props.value]);

  useEffect(() => {
    if (typeof props.onValueChange === "function") {
      props.onValueChange(dialogOpen.value);
    }
  }, [dialogOpen.value]);

  function dialogClick(event: MouseEvent) {
    if (
      event.target &&
      (event.target as HTMLElement).tagName === "DIALOG" &&
      closeOnClick.value
    ) {
      dialogOpen.value = false;
    }
  }

  return (
    <>
      {!props.hideActivator && (
        <button onClick={() => (dialogOpen.value = true)}>Open</button>
      )}
      <dialog
        open={dialogOpen.value}
        class="fixed top-0 left-0 m-0 w-full h-full bg-black bg-opacity-20"
        onClick={dialogClick}
      >
        <div class="bg-white w-full max-w-lg mx-auto my-8 p-2 rounded">
          {props.title && <p class="text-xl mb-2">{props.title}</p>}
          <div>{props.children}</div>
          <div class="flex justify-end items-center mt-2">
            {props.actionButtons}
            <button
              class="py-1 px-2 rounded border bg-gray-200"
              onClick={() => (dialogOpen.value = false)}
            >
              Close
            </button>
          </div>
        </div>
      </dialog>
    </>
  );
}
