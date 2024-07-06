
export const HeaderBar = () => {
    return (
        <div className="flex justify-between items-center bg-gray-800 text-white p-4">

            <div className="flex items-center">
                <img src="/path-to-your-logo.png" alt="Logo" className="h-8" />
            </div>


            <div className="text-sm">
                <a href="#" className="text-gray-300 hover:text-white mr-4">Mi Perfil</a>
            </div>


            <div className="flex items-center">

                <button className="bg-gray-700 hover:bg-gray-600 text-white py-2 px-4 rounded mr-4">
                    Currency
                </button>


                <button className="bg-red-600 hover:bg-red-500 text-white py-2 px-4 rounded">
                    Cerrar Sesión
                </button>
            </div>
        </div>
    );
};


