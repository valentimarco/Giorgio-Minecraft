import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react";
import type { LinksFunction } from "@remix-run/node";
import { Icon } from '@iconify/react';

import "./tailwind.css";
import { useState } from "react";
import { ServerMenu } from "./components/layout/serverMenu";

export const links: LinksFunction = () => [
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  {
    rel: "preconnect",
    href: "https://fonts.gstatic.com",
    crossOrigin: "anonymous",
  },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  },
];

export function Layout({ children }: { children: React.ReactNode }) {
  const [showNav, setShowNav] = useState<boolean>(true)

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="flex">
        <div className={`flex flex-col bg-slate-500 h-screen ${showNav ? 'w-[256px]' : 'w-0'}`}>
          <div className={`flex justify-between`}>
            <h1 className={`text-center flex-grow ${showNav ? '' : 'hidden'}`}>List Servers</h1>
            <button className="rounded-lg" onClick={() => setShowNav(!showNav)}><Icon icon="octicon:three-bars-24" width="24" height="24" /></button>
          </div>
          <nav className="flex-grow overflow-y-auto">
            {
              Array.from({ length: 2 }, (_, index) => (
                <ServerMenu key={index} serverName={"Hello"} serverId={"ops"} />
              ))

            }
          </nav>
          <div className={`flex justify-center py-6 ${showNav ? '' : 'hidden'}`}>
            <h1 className="text-sm">User 1</h1>
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
