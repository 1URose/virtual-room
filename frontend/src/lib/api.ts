import { Event } from '@/types/event'

const API_URL = '/api/events'

export async function fetchEvents(): Promise<Event[]> {
    const response = await fetch(`${API_URL}/all`)
    if (!response.ok) {
        throw new Error('Failed to fetch events')
    }
    return response.json()
}

export async function createEvent(event: Event): Promise<Event> {
    const response = await fetch(`${API_URL}/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(event),
    })
    if (!response.ok) {
        throw new Error('Failed to create event')
    }
    return response.json()
}

export async function updateEvent(id: number, event: Event): Promise<Event> {
    const response = await fetch(`${API_URL}/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(event),
    })
    if (!response.ok) {
        throw new Error('Failed to update event')
    }
    return response.json()
}

export async function deleteEvent(id: number): Promise<void> {
    const response = await fetch(`${API_URL}/${id}`, {
        method: 'DELETE',
    })
    if (!response.ok) {
        throw new Error('Failed to delete event')
    }
}

