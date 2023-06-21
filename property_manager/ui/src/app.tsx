import { ConnectionHeader } from "./components/ConnectionHeader";
import { PropertyList } from "./components/PropertyList";
import { AppState, createState } from "./state";

export function App() {
  return (
    <AppState.Provider value={createState()}>
      <div className="flex flex-col max-w-2xl mx-auto">
        <ConnectionHeader />
        <hr />
        <PropertyList />
      </div>
    </AppState.Provider>
  );
}
