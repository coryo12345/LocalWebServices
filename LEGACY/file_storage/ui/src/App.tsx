import { ConenctionControl } from "./components/ConnectionControl";
import { FileStorage } from "./components/FileStorage";
import { GlobalProgress } from "./components/GlobalProgress";
import { ThemeToggle } from "./components/theme/ThemeToggle";
import { Toaster } from "./components/ui/toaster";

export function App() {
  return (
    <>
      <GlobalProgress />
      <main className="max-w-4xl mx-auto mt-6">
        <div className="flex justify-between mb-2">
          <h1 className="text-xl inline-block">File Storage</h1>
          <ThemeToggle className="inline-block" />
        </div>
        <ConenctionControl className="mb-2" />
        <FileStorage />
      </main>
      <Toaster />
    </>
  );
}

export default App;
