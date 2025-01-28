import { Icon } from '@iconify/react';
import { DragEvent, useState } from 'react';



export function CreateServerModal() {
    const [listSwitch, setListSwitch] = useState<string>("custom")


    const handleDrop = (e: DragEvent<HTMLDivElement>) => {
        e.preventDefault();
        // Handle file drop here
        const file = e.dataTransfer?.files[0];
        console.log(file)
    };

    const renderContent = () => {
        switch (listSwitch) {
            case "custom":
                return (
                    <div className='flex flex-col justify-center'>
                        <h1 className='text-center text-lg'>Drag and Drop your minecraft server</h1>
                        <h1 className='text-center text-lg'>Supported List:</h1>
                        <ol type='i'>
                            <li>
                                <h1 className='text-center text-lg'>CurseForge Zips</h1>
                            </li>
                            <li>
                                <h1 className='text-center text-lg'>FTB Zips</h1>
                            </li>
                        </ol>
                        <div
                            onDrop={handleDrop}
                            className="my-12 border-2 border-dashed rounded-lg p-12 text-center transition-colors 
                            cursor-pointer border-base-300 hover:border-primary hover:bg-primary/10"
                        >
                            <p className="text-lg">Drag and drop your Zip here</p>
                            <p className="text-sm text-base-content/60 mt-2">
                                Only Zip files are accepted
                            </p>
                        </div>
                    </div>
                )
            case "CF":
                return (
                    <div className='flex flex-col justify-center'>
                        <h1 className='text-center text-lg'>Work in progress</h1>
                    </div>
                )
            case "FTB":
                return (
                    <div className='flex flex-col justify-center'>
                        <h1 className='text-center text-lg'>Work in progress</h1>
                    </div>
                )

        }

    }

    return (
        <>
            <button className="btn btn-secondary btn-sm" onClick={() => (document.getElementById('create_server') as HTMLDialogElement).showModal()}>
                <Icon icon="octicon:plus-24" />
                <span className="text-sm">Add Server</span>
            </button>
            <dialog id="create_server" className="modal">
                <div className="modal-box max-w-7xl h-5/6 w-11/12">

                    <div className="grid grid-cols-[12rem_1fr] h-full">
                        <div className='flex flex-col m-2 items-center border-r'>
                            <h1 className='text-lg'>Provider</h1>
                            <ul className="menu grow bg-base-200 rounded-box m-2 w-10/12">
                                <li><button onClick={() => setListSwitch("custom")}><p className='text-center text-lg'>Custom</p></button></li>
                                <li><button onClick={() => setListSwitch("CF")}><p className='text-center text-lg'>CurseForge</p></button></li>
                                <li><button onClick={() => setListSwitch("FTB")}><p className='text-center text-lg'>FTB</p></button></li>
                            </ul>
                        </div>

                        {renderContent()}

                    </div>




                    <form method="dialog">
                        <button className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
                            <Icon icon="octicon:x-24" width="24" height="24" />
                        </button>
                    </form>
                </div>
            </dialog>
        </>
    );
}