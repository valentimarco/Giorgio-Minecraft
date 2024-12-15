export default function ChartSkeleton() {
    return (
        <div className="flex flex-col">
            <div className="skeleton h-8 w-48"></div>


            <div className="relative h-64">

                <div className="absolute left-0 top-0 h-full flex flex-col justify-between">
                    {[...Array(5)].map((_, i) => (
                        <div key={i} className="skeleton h-4 w-8"></div>
                    ))}
                </div>


                <div className="ml-12 h-full flex items-end justify-between gap-4">
                    <div className="skeleton w-12 h-3/4"></div>
                    <div className="skeleton w-12 h-1/2"></div>
                    <div className="skeleton w-12 h-full"></div>
                    <div className="skeleton w-12 h-2/3"></div>
                    <div className="skeleton w-12 h-4/5"></div>
                    <div className="skeleton w-12 h-1/3"></div>
                    <div className="skeleton w-12 h-3/5"></div>
                </div>


                <div className="absolute bottom--8 left-12 right-0 flex justify-between">
                    {[...Array(7)].map((_, i) => (
                        <div key={i} className="skeleton h-4 w-12"></div>
                    ))}
                </div>
            </div>
        </div>
    );
}