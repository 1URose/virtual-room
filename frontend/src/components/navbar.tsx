import Link from 'next/link'

export function Navigation() {
    return (
        <nav className="bg-gray-800 p-4">
            <div className="container mx-auto flex justify-between items-center">
                <Link href="/" className="text-xl font-bold text-primary">Event Management</Link>
                <div className="space-x-4">
                    <Link href="/events" className="text-gray-300 hover:text-primary">Events</Link>
                    <Link href="/login" className="text-gray-300 hover:text-primary">Login</Link>
                    <Link href="/signup" className="text-gray-300 hover:text-primary">Register</Link>
                </div>
            </div>
        </nav>
    )
}

