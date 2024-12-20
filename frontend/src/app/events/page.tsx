'use client'

import { useState, useEffect } from 'react'

// Тип события, которое уже есть в базе (подходящее для отображения)
interface Event {
    id: number
    event_name: string
    description: string
    start_time: string
    end_time: string
    // Можно добавить поля, если детальный эндпоинт возвращает больше данных
}

export default function Events() {
    // Состояние для списка событий
    const [events, setEvents] = useState<Event[]>([])
    // Флаги загрузки и ошибок
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)

    // Поля для создания события (соответствуют CreateEvent)
    const [newEventTitle, setNewEventTitle] = useState('')
    const [newEventDescription, setNewEventDescription] = useState('')
    const [newVirtualRoomName, setNewVirtualRoomName] = useState('default_room')
    const [newOrganizerLogin, setNewOrganizerLogin] = useState('demo_organizer')
    const [newEventStartTime, setNewEventStartTime] = useState('')
    const [newEventEndTime, setNewEventEndTime] = useState('')

    // Детальный просмотр
    const [selectedEvent, setSelectedEvent] = useState<Event | null>(null)
    // Флаг загрузки для детального просмотра (если нужно)
    const [detailLoading, setDetailLoading] = useState(false)

    // При монтировании компонента — загрузить список событий
    useEffect(() => {
        fetchEvents()
    }, [])

    // Функция получения всех событий (GET /events/all)
    async function fetchEvents() {
        try {
            const response = await fetch('http://localhost:8080/events/all')
            if (!response.ok) {
                throw new Error('Failed to fetch events')
            }
            const data = await response.json()
            setEvents(data.events || [])
        } catch (err) {
            setError('An error occurred while fetching events')
        } finally {
            setLoading(false)
        }
    }

    // Функция повторного запроса всех событий
    async function fetchAllEvents() {
        setLoading(true)
        setError(null)
        try {
            const response = await fetch('http://localhost:8080/events/all')
            if (!response.ok) {
                throw new Error('Failed to fetch events')
            }
            const data = await response.json()
            console.log(data.events)
            setEvents(data.events || [])
        } catch (err) {
            setError('An error occurred while fetching events')
        } finally {
            setLoading(false)
        }
    }

    // Функция создания события (POST /events/create)
    async function createEvent() {
        try {
            // Преобразуем datetime-local в строку ISO-8601 (RFC3339)
            const startDateISO = new Date(newEventStartTime).toISOString()
            const endDateISO = new Date(newEventEndTime).toISOString()
            console.log(startDateISO)

            // Тело запроса должно соответствовать структуре CreateEvent
            const requestBody = {
                title: newEventTitle.trim(),
                description: newEventDescription.trim(),
                virtual_room_name: newVirtualRoomName.trim(),
                organizer_login: newOrganizerLogin.trim(),
                start_date: startDateISO,  // Будет распарсено как time.Time
                end_date: endDateISO       // Аналогично
            }

            const response = await fetch('http://localhost:8080/events/create', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(requestBody),
            })
            if (!response.ok) {
                throw new Error('Failed to create event')
            }

            const newData = await response.json()
            console.log('Создано событие:', newData)

            // Обновляем список событий
            await fetchAllEvents()

            // Очищаем форму
            clearForm()
        } catch (err) {
            setError('An error occurred while creating the event')
            console.error(err)
        }
    }

    // Получить детальные данные события (GET /events/{id})
    async function handleViewEvent(id: number) {
        try {
            setDetailLoading(true)
            const response = await fetch(`http://localhost:8080/events/${id}`)
            if (!response.ok) {
                throw new Error('Failed to fetch event details')
            }
            const eventData: Event = await response.json()
            setSelectedEvent(eventData)
        } catch (error) {
            console.error('Error fetching event details:', error)
        } finally {
            setDetailLoading(false)
        }
    }

    // Очистить поля формы
    function clearForm() {
        setNewEventTitle('')
        setNewEventDescription('')
        setNewVirtualRoomName('default_room')
        setNewOrganizerLogin('demo_organizer')
        setNewEventStartTime('')
        setNewEventEndTime('')
    }

    if (loading) {
        return <div className="flex justify-center items-center min-h-screen">Loading...</div>
    }

    if (error) {
        return <div className="flex justify-center items-center min-h-screen text-red-500">{error}</div>
    }

    return (
        <div className="container mx-auto px-4 py-8">
            <h1 className="text-3xl font-bold mb-8 text-center">Upcoming Events</h1>

            {/* Форма для создания нового события */}
            <div className="max-w-xl mx-auto mb-10 bg-gray-800 p-6 rounded-lg shadow-lg">
                <h2 className="text-xl font-semibold mb-4 text-primary">Create New Event</h2>
                <form
                    onSubmit={(e) => {
                        e.preventDefault()
                        createEvent()
                    }}
                    className="space-y-4"
                >
                    <div>
                        <label className="block mb-1 text-gray-300">Title</label>
                        <input
                            type="text"
                            value={newEventTitle}
                            onChange={(e) => setNewEventTitle(e.target.value)}
                            className="w-full rounded px-3 py-2 bg-gray-900 text-gray-100"
                            placeholder="Enter event title"
                            required
                        />
                    </div>
                    <div>
                        <label className="block mb-1 text-gray-300">Description</label>
                        <textarea
                            value={newEventDescription}
                            onChange={(e) => setNewEventDescription(e.target.value)}
                            className="w-full rounded px-3 py-2 bg-gray-900 text-gray-100"
                            placeholder="Enter event description"
                            rows={3}
                        />
                    </div>
                    <div>
                        <label className="block mb-1 text-gray-300">Virtual Room Name</label>
                        <input
                            type="text"
                            value={newVirtualRoomName}
                            onChange={(e) => setNewVirtualRoomName(e.target.value)}
                            className="w-full rounded px-3 py-2 bg-gray-900 text-gray-100"
                            placeholder="Enter virtual room name"
                            required
                        />
                    </div>
                    <div>
                        <label className="block mb-1 text-gray-300">Organizer Login</label>
                        <input
                            type="text"
                            value={newOrganizerLogin}
                            onChange={(e) => setNewOrganizerLogin(e.target.value)}
                            className="w-full rounded px-3 py-2 bg-gray-900 text-gray-100"
                            placeholder="Enter organizer login"
                            required
                        />
                    </div>
                    <div>
                        <label className="block mb-1 text-gray-300">Start Date</label>
                        <input
                            type="datetime-local"
                            value={newEventStartTime}
                            onChange={(e) => setNewEventStartTime(e.target.value)}
                            className="w-full rounded px-3 py-2 bg-gray-900 text-gray-100"
                            required
                        />
                    </div>
                    <div>
                        <label className="block mb-1 text-gray-300">End Date</label>
                        <input
                            type="datetime-local"
                            value={newEventEndTime}
                            onChange={(e) => setNewEventEndTime(e.target.value)}
                            className="w-full rounded px-3 py-2 bg-gray-900 text-gray-100"
                            required
                        />
                    </div>

                    <button
                        type="submit"
                        className="bg-primary text-primary-foreground px-4 py-2 rounded hover:bg-primary/90 transition-colors"
                    >
                        Create
                    </button>
                </form>
            </div>

            {/* Список событий */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {events.length > 0 && events.map((event) => (
                    <div key={event.id} className="bg-gray-800 rounded-lg shadow-lg overflow-hidden">
                        <div className="p-6">
                            <h2 className="text-xl font-semibold mb-2 text-primary">{event.event_name}</h2>
                            <p className="text-gray-300 mb-4">{event.description}</p>
                            <div className="text-sm text-gray-400">
                                <p>Start: {new Date(event.start_time).toLocaleString()}</p>
                                <p>End: {new Date(event.end_time).toLocaleString()}</p>
                            </div>
                        </div>
                        <div className="px-6 py-4 bg-gray-900 flex gap-4">
                            {/* Кнопка для просмотра детальной информации на этой же странице */}
                            <button
                                onClick={() => handleViewEvent(event.id)}
                                className="text-primary hover:text-primary-foreground transition-colors underline"
                            >
                                View Details
                            </button>
                            {/* Или ссылка на отдельную страницу /events/[id] (если есть такая) */}
                            {/* <Link
                href={`/events/${event.id}`}
                className="text-primary hover:text-primary-foreground transition-colors"
              >
                Go to page
              </Link> */}
                        </div>
                    </div>
                ))}
            </div>

            {/* Детальный просмотр события на этой же странице */}
            {selectedEvent && (
                <div className="max-w-xl mx-auto mt-10 bg-gray-800 p-6 rounded-lg shadow-lg">
                    {detailLoading ? (
                        <p className="text-gray-400">Loading detail...</p>
                    ) : (
                        <>
                            <h3 className="text-xl font-semibold mb-4 text-primary">
                                Event ID: {selectedEvent.id}
                            </h3>
                            <p className="text-gray-300 mb-2">
                                Event Name: {selectedEvent.event_name}
                            </p>
                            <p className="text-gray-300 mb-2">
                                Description: {selectedEvent.description}
                            </p>
                            <p className="text-sm text-gray-400">
                                Start: {new Date(selectedEvent.start_time).toLocaleString()}
                            </p>
                            <p className="text-sm text-gray-400">
                                End: {new Date(selectedEvent.end_time).toLocaleString()}
                            </p>

                            <button
                                onClick={() => setSelectedEvent(null)}
                                className="mt-4 bg-gray-700 hover:bg-gray-600 text-gray-100 px-4 py-2 rounded"
                            >
                                Close
                            </button>
                        </>
                    )}
                </div>
            )}
        </div>
    )
}