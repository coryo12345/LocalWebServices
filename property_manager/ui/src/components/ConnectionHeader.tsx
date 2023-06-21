import { useContext } from "preact/hooks";
import { AppState } from "../state";

export function ConnectionHeader() {
  const state = useContext(AppState);

  function onSubmit(e: Event) {
    e.preventDefault();
    state.fetchProperties();
  }

  return (
    <div className="m-2">
      <form onSubmit={onSubmit}>
        <input
          type="text"
          placeholder="http://localhost:8081"
          className="py-1 px-2 border border-black rounded mr-2"
          onInput={(e) => (state.url.value = e.currentTarget.value)}
        />
        <button
          type="submit"
          className="py-1 px-2 rounded border bg-blue-700 text-white"
        >
          Submit
        </button>
      </form>
    </div>
  );
}
