import type { MetaFunction } from "@remix-run/node";



export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  return (
    <div className="grid grid-cols-2 ml-5 h-screen w-screen overflow-auto">

    </div>
  );
}

