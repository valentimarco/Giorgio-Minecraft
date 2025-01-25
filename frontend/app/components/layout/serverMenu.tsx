import { Link } from "@remix-run/react";

export type ServerMenuProps = {
    serverName: string,
    serverId: string
}

export function ServerMenu({ serverName, serverId }: ServerMenuProps) {
    return (
        <div className="w-full my-2">
            <details className="w-full">
                <summary className="btn btn-primary w-full list-none cursor-pointer">
                    {serverName}
                </summary>
                <ul className="menu w-full bg-base-200 rounded-lg">
                    <li><Link to={`${serverId}/dashboard`}>Dashboard</Link></li>
                    <li><Link to={`${serverId}/terminal`}>Terminal</Link></li>
                    <li><Link to={`${serverId}/config`}>Config</Link></li>
                </ul>
            </details>
        </div>
    )
}