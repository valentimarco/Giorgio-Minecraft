import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
} from "@remix-run/react";
import type { LinksFunction } from "@remix-run/node";
import "./tailwind.css";

import { ServerInfo } from "./types";
import { CreateServerModal, ServerMenu, ServerHeader } from "./components"


export const links: LinksFunction = () => [
  // { rel: "preconnect", href: "https://fonts.googleapis.com" },
  // {
  //   rel: "preconnect",
  //   href: "https://fonts.gstatic.com",
  //   crossOrigin: "anonymous",
  // },
  // {
  //   rel: "stylesheet",
  //   href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  // },
];


export async function loader(): Promise<ServerInfo[]> {
  //GET /api/v1/servers?query=name,id

  return [
    {
      name: "hello",
      id: "ops"
    },
    {
      name: "world",
      id: "spo"
    }
  ]
}

export function Layout({ children }: { children: React.ReactNode }) {
  const servers = useLoaderData<typeof loader>();

  return (
    <html lang="en" data-theme="dark">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="flex min-h-dvh scroll-smooth antialiased transition-colors bg-base-100">
        <div className="grid grid-cols-[16rem_1fr] grow">
          <div className={`flex flex-col bg-base-100 rounded-lg m-2 border-2`}>

            <div className={`flex justify-center p-4 border-b border-slate-700`}>
              <h1 className={`text-slate-100 font-semibold text-lg`}>List Servers</h1>
              <CreateServerModal />
            </div>

            <nav className="grow overflow-y-auto px-3 py-4">
              {/* TODO: pass the server obj instead of values */}
              {
                servers.map((server, index) => (
                  <ServerMenu key={index} serverName={server.name} serverId={server.id} />
                ))
              }
            </nav>

            <div className={`border-t border-slate-700 p-4`}>
              <div className="flex items-center gap-3 text-slate-300">
                <div className="w-8 h-8 bg-slate-700 rounded-full flex items-center justify-center">
                  <h1 className="text-sm">Icon</h1>
                </div>
                <span className="text-sm font-medium">User 1</span>
              </div>
            </div>

          </div>

          <div className="grid grid-rows-[4rem_1fr] m-2 border">
            <ServerHeader />

            {children}
          </div>
        </div>

        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return <Outlet />;
}
