import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { useState, useRef, useEffect } from "react";



export async function loader({ params }: LoaderFunctionArgs) {

    //GET /api/v1/servers/{id}/info




    return {
        serverName: "hello",
        serverId: params.serverId
    }
}

export default function Servers() {
    const { serverId, serverName } = useLoaderData<typeof loader>();

    const [logs, setLogs] = useState([
        { timestamp: new Date().toLocaleTimeString(), message: "Server started successfully", type: "info" },
        { timestamp: new Date().toLocaleTimeString(), message: "Waiting for players...", type: "info" }
    ]);
    const [command, setCommand] = useState('');
    const logsRef = useRef<HTMLDivElement>(null);

    // Auto scroll to bottom when new logs arrive
    useEffect(() => {
        if (logsRef.current) {
            logsRef.current.scrollTop = logsRef.current.scrollHeight;
        }
    }, [logs]);

    const handleCommand = (e: { preventDefault: () => void; }) => {
        e.preventDefault();
        if (!command.trim()) return;

        // Simulate command processing
        setLogs(prev => [...prev, {
            timestamp: new Date().toLocaleTimeString(),
            message: `> ${command}`,
            type: "command"
        }]);

        // Here you would typically send the command to your server
        // For now, we'll just simulate a response
        setTimeout(() => {
            setLogs(prev => [...prev, {
                timestamp: new Date().toLocaleTimeString(),
                message: `Executed command: ${command}`,
                type: "response"
            }]);
        }, 100);

        setCommand('');
    };
    return (
        <div className="flex-1 flex flex-col bg-slate-900 text-slate-100">
            {/* Server Info Header */}
            <div className="border-b border-slate-700 p-4">
                <div className="flex items-center justify-between">
                    <h2 className="text-lg font-semibold text-slate-100"> {serverName} Console</h2>
                    <div className="flex items-center space-x-2">
                        <div className="w-3 h-3 rounded-full bg-green-500"></div>
                        <span className="text-sm text-slate-300">Online</span>
                    </div>
                </div>
            </div>

            {/* Logs Display */}
            <div
                ref={logsRef}
                className="flex-auto grow overflow-y-auto p-4 font-mono text-sm bg-slate-950"
            >
                {logs.map((log, index) => (
                    <div key={index} className="mb-1">
                        <span className="text-slate-500">[{log.timestamp}] </span>
                        <span className={`
                            ${log.type === 'error' && 'text-red-400'}
                            ${log.type === 'info' && 'text-blue-400'}
                            ${log.type === 'command' && 'text-yellow-400'}
                            ${log.type === 'response' && 'text-green-400'}
                        `}>
                            {log.message}
                        </span>
                    </div>
                ))}
            </div>

            {/* Command Input */}
            <div className=" border-slate-700 p-3">
                <form onSubmit={handleCommand} className="flex space-x-2">
                    <div className="flex-1 relative">
                        <input
                            type="text"
                            value={command}
                            onChange={(e) => setCommand(e.target.value)}
                            className="w-full bg-slate-800 text-slate-100 px-4 py-2 rounded-lg border border-slate-700 focus:outline-none focus:border-slate-500"
                            placeholder="Type command here..."
                        />
                    </div>
                    <button
                        type="submit"
                        className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
                    >
                        Send
                    </button>
                </form>
            </div>
        </div>
    );
}