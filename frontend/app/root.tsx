import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
} from "@remix-run/react";
import type { LinksFunction } from "@remix-run/node";
import { Icon } from '@iconify/react';

import "./tailwind.css";
import { useState } from "react";
import { ServerMenu } from "./components/layout/serverMenu";
import { ServerInfo } from "./types";

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
  const [showNav, setShowNav] = useState<boolean>(true)
  const serverInfo = useLoaderData<typeof loader>();

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="flex overflow-hidden">
        <div className={`flex flex-col bg-gradient-to-b from-slate-800 to-slate-900 h-screen transition-all duration-300 ease-in-out ${showNav ? 'w-64' : 'w-0'}`}>
          <div className={`flex justify-between items-center p-4 border-b border-slate-700`}>
            <h1 className={`text-slate-100 font-semibold text-lg ${showNav ? '' : 'hidden'}`}>List Servers</h1>
            <button className="rounded-lg" onClick={() => setShowNav(!showNav)}><Icon icon="octicon:three-bars-24" width="24" height="24" /></button>
          </div>
          <nav className="flex-grow overflow-y-auto px-3 py-4">
            {
              serverInfo.map((server, index) => (
                <ServerMenu key={index} serverName={server.name} serverId={server.id} />
              ))
            }
          </nav>

          <div className={`border-t border-slate-700 p-4 ${showNav ? '' : 'hidden'}`}>
            <div className="flex items-center gap-3 text-slate-300">
              <div className="w-8 h-8 bg-slate-700 rounded-full flex items-center justify-center">
                <h1 className="text-sm">Icon</h1>
              </div>
              <span className="text-sm font-medium">User 1</span>
            </div>
          </div>
        </div>
        {children}
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return <Outlet />;
}
