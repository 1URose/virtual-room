import Link from 'next/link'

export default function Home() {
  return (
      <main className="flex min-h-screen flex-col items-center justify-center p-24">
        <h1 className="text-4xl font-bold mb-8">Event Management System</h1>
        <div className="flex gap-4">
          <Link href="/login" className="bg-primary text-primary-foreground px-4 py-2 rounded-md hover:bg-primary/90">
            Login
          </Link>
          <Link href="/signup" className="bg-secondary text-secondary-foreground px-4 py-2 rounded-md hover:bg-secondary/90">
            Register
          </Link>
        </div>
      </main>
  )
}

