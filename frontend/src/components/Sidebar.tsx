'use client'

import Link from 'next/link'
import {Briefcase, Calendar, Monitor, PenTool, Ticket} from 'lucide-react'
import {
    Sidebar,
    SidebarContent,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
} from '@/components/ui/sidebar'

const menuItems = [
    { icon: Calendar, label: 'Events', href: '/events' },
    { icon: Ticket, label: 'Tickets', href: '/tickets' },
    { icon: Briefcase, label: 'Sponsors', href: '/sponsors' },
    { icon: Monitor, label: 'Virtual Rooms', href: '/virtual-rooms' },
    { icon: PenTool, label: 'Equipment', href: '/equipment' },
]

export function AppSidebar() {
    return (
        <Sidebar className="border-r border-border/50 bg-gray-900 text-white">
            <SidebarHeader className="border-b border-border/50 p-4">
                <Link href="/" className="flex items-center gap-2 text-xl font-semibold text-primary">
                    <Calendar className="h-6 w-6" />
                    Event Management
                </Link>
            </SidebarHeader>
            <SidebarContent>
                <SidebarMenu>
                    {menuItems.map((item) => (
                        <SidebarMenuItem key={item.href}>
                            <SidebarMenuButton asChild>
                                <Link
                                    href={item.href}
                                    className="flex items-center gap-3 px-3 py-2 text-sm hover:text-primary"
                                >
                                    <item.icon className="h-4 w-4" />
                                    <span>{item.label}</span>
                                </Link>
                            </SidebarMenuButton>
                        </SidebarMenuItem>
                    ))}
                </SidebarMenu>
            </SidebarContent>
        </Sidebar>
    )
}