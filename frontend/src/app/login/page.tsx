'use client'

import {useState} from 'react'
import {useRouter} from 'next/navigation'

export default function Login() {
    const [login, setLogin] = useState('')
    const [password, setPassword] = useState('')
    const router = useRouter()

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await fetch('http://localhost:8080/users/auth', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({login, password}),
            })
            if (response.ok) {
                const newVar = await response.json();
                localStorage.setItem("token", JSON.stringify(newVar.token))
                router.push('/events')
            } else {
                // Handle error
                console.error('Login failed')
            }
        } catch (error) {
            console.error('An error occurred', error)
        }
    }

    return (
        <div className="flex min-h-screen flex-col items-center justify-center p-24">
            <h1 className="text-3xl font-bold mb-8">Login</h1>
            <form onSubmit={handleSubmit} className="w-full max-w-xs">
                <input
                    type="text"
                    value={login}
                    onChange={(e) => setLogin(e.target.value)}
                    placeholder="Login"
                    className="w-full p-2 mb-4 rounded bg-gray-800 text-white"
                    required
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Password"
                    className="w-full p-2 mb-4 rounded bg-gray-800 text-white"
                    required
                />
                <button type="submit"
                        className="w-full bg-primary text-primary-foreground p-2 rounded hover:bg-primary/90">
                    Login
                </button>
            </form>
        </div>
    )
}

