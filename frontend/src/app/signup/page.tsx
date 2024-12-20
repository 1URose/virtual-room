'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'

export default function Register() {
    const [name, setName] = useState('')
    const [login, setLogin] = useState('')
    const [password, setPassword] = useState('')
    const [userRole, setUserRole] = useState(1)
    const router = useRouter()

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await fetch('http://localhost:8080/users/signUp', {
                method: 'POST',
                body: JSON.stringify({ name, login, password, user_role: userRole }),
            })
            if (response.ok) {
                router.push('/login')
            } else {
                // Handle error
                console.error('Registration failed')
            }
        } catch (error) {
            console.error('An error occurred', error)
        }
    }

    return (
        <div className="flex min-h-screen flex-col items-center justify-center p-24">
            <h1 className="text-3xl font-bold mb-8">Register</h1>
            <form onSubmit={handleSubmit} className="w-full max-w-xs">
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="Name"
                    className="w-full p-2 mb-4 rounded bg-gray-800 text-white"
                    required
                />
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
                <select
                    value={userRole}
                    onChange={(e) => setUserRole(parseInt(e.target.value))}
                    className="w-full p-2 mb-4 rounded bg-gray-800 text-white"
                >
                    <option value="1">Participant</option>
                    <option value="2">Organizer</option>
                    <option value="3">Admin</option>
                </select>
                <button type="submit" className="w-full bg-primary text-primary-foreground p-2 rounded hover:bg-primary/90">
                    Register
                </button>
            </form>
        </div>
    )
}

