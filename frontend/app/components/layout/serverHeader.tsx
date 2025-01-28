import { Icon } from '@iconify/react';

export function ServerHeader() {
    return (
        <div className="flex items-center">
            <button className="btn btn-neutral btn-sm rounded-lg m-2"><Icon icon="octicon:three-bars-24" width="12" height="12" /></button>
            <div className='flex grow justify-center'>
                <h1 className='text-center text-lg'>ServerName</h1>
            </div>
            <div className="badge badge-success m-2" >Online</div>
        </div>
    );
}