import { ConnectionHeader } from "./components/ConnectionHeader";
import { PropertyList } from "./components/PropertyList";
import { Snackbar } from "./components/base/Snackbar";
import { AppState, createState } from "./state";

export function App() {
  return (
    <AppState.Provider value={createState()}>
      <div className="flex flex-col max-w-2xl mx-auto">
        <ConnectionHeader />
        <hr />
        <PropertyList />
      </div>
      <Snackbar />
    </AppState.Provider>
  );
}
