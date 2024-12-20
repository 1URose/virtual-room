"use client"

import * as React from "react"
import { cn } from "@/lib/utils"

const SidebarContext = React.createContext<{
    expanded: boolean
    setExpanded: (expanded: boolean) => void
} | undefined>(undefined)

export function Sidebar({ className, children }: React.HTMLAttributes<HTMLDivElement>) {
    const [expanded, setExpanded] = React.useState(true)

    return (
        <SidebarContext.Provider value={{ expanded, setExpanded }}>
            <div className={cn("flex h-screen", className)}>
                <div
                    className={cn(
                        "flex flex-col bg-muted text-muted-foreground",
                        expanded ? "w-64" : "w-16"
                    )}
                >
                    {children}
                </div>
            </div>
        </SidebarContext.Provider>
    )
}

export function SidebarHeader({ className, children }: React.HTMLAttributes<HTMLDivElement>) {
    return <div className={cn("p-4", className)}>{children}</div>
}

export function SidebarContent({ className, children }: React.HTMLAttributes<HTMLDivElement>) {
    return <div className={cn("flex-1 overflow-auto", className)}>{children}</div>
}

export function SidebarMenu({ className, children }: React.HTMLAttributes<HTMLUListElement>) {
    return <ul className={cn("space-y-2", className)}>{children}</ul>
}

export function SidebarMenuItem({ className, children }: React.HTMLAttributes<HTMLLIElement>) {
    return <li className={cn("px-4", className)}>{children}</li>
}

export function SidebarMenuButton({
                                      className,
                                      children,
                                      ...props
                                  }: React.ButtonHTMLAttributes<HTMLButtonElement>) {
    const { expanded } = React.useContext(SidebarContext) || { expanded: true }

    return (
        <button
            className={cn(
                "flex items-center w-full p-2 rounded-md hover:bg-accent hover:text-accent-foreground",
                expanded ? "justify-start" : "justify-center",
                className
            )}
            {...props}
        >
            {children}
        </button>
    )
}

export function SidebarProvider({ children }: { children: React.ReactNode }) {
    return <>{children}</>
}

