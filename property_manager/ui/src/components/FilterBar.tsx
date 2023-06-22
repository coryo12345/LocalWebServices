import { Menu, MenuButton, MenuItem } from "@szhsin/react-menu";
import { useContext } from "preact/hooks";
import { AppState } from "../state";
import { ListAction } from "./base/DataTable";

interface Props {
  onAction?: (action: ListAction) => void;
}

export function FilterBar(props: Props) {
  const state = useContext(AppState);

  function deleteAction() {
    typeof props.onAction === "function" && props.onAction("delete");
  }

  function addAction() {
    typeof props.onAction === "function" && props.onAction("add");
  }

  return (
    <div class="m-2 flex items-center">
      <input
        type="text"
        placeholder="Search for a property"
        className="py-1 px-2 border border-black rounded mr-2"
        onInput={(e) => (state.searchText.value = e.currentTarget.value)}
      />
      <span className="mr-2">
        <input
          checked={state.searchKey}
          id="search-key-checkbox"
          type="checkbox"
          className="mr-1 h-4 w-4"
          onChange={() => (state.searchKey.value = !state.searchKey.value)}
        />
        <label for="search-key-checkbox" className="h-full">
          Search Key
        </label>
      </span>
      <span>
        <input
          checked={state.searchValue}
          id="search-value-checkbox"
          type="checkbox"
          className="mr-1 h-4 w-4"
          onChange={() => (state.searchValue.value = !state.searchValue.value)}
        />
        <label for="search-value-checkbox">Search Value</label>
      </span>
      <span className="ml-auto">
        <Menu menuButton={MB} transition>
          <MenuItem onClick={() => addAction()}>Add New Property</MenuItem>
          <MenuItem onClick={() => deleteAction()}>Bulk Delete</MenuItem>
        </Menu>
      </span>
    </div>
  );
}

const MB = (
  <MenuButton className="border border-slate-500 rounded py-1 px-2">
    Action &#9660;
  </MenuButton>
);
