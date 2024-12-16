import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";



export async function loader({ params }: LoaderFunctionArgs) {
    return {
        serverId: params.serverId
    }
}

export default function Servers() {
    const data = useLoaderData<typeof loader>();
    return (
        <div className="flex justify-center">
            <p>{data.serverId}</p>
        </div>
    );
}