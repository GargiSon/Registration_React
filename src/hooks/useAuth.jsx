import { useState, useEffect } from "react";

export function useAuth(){
    const [isAuthenticated, setIsAuthenticated] = useState(null);

    useEffect(() => {
        const token = localStorage.getItem('token');
        setIsAuthenticated(!!token);
    }, []);

    return isAuthenticated;
}