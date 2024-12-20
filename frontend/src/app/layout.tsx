import { Inter } from 'next/font/google'
import './globals.css'
import { SidebarProvider } from '@/components/ui/sidebar'
import {AppSidebar} from "@/components/Sidebar";

const inter = Inter({ subsets: ['latin'] })

export const metadata = {
    title: 'Event Management System',
    description: 'Manage events, tickets, and more',
}

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (
        <html lang="en" className="dark">
        <body className={inter.className}>
        <SidebarProvider>
            <div className="flex min-h-screen bg-purple-900">
                <AppSidebar />
                <main className="flex-1 p-8">
                    {children}
                </main>
            </div>
        </SidebarProvider>
        </body>
        </html>
    )
}

