'use client';

import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';

type TTicket = {
    login: string;
    event_name: string;
    price: number;
    ticket_type: number; // 0 = Free, 1 = Paid
};

export default function TicketsPage() {
    // Состояние для списка билетов
    const [tickets, setTickets] = useState<TTicket[]>([]);

    // Поля формы для создания/редактирования
    const [login, setLogin] = useState('');
    const [eventName, setEventName] = useState('');
    const [price, setPrice] = useState('');
    const [ticketType, setTicketType] = useState<number>(0); // по умолчанию Free (0)

    // Режим редактирования
    const [isEditMode, setIsEditMode] = useState(false);
    // Оригинальные поля билета (для поиска в PUT/DELETE)
    const [originalLogin, setOriginalLogin] = useState('');
    const [originalEventName, setOriginalEventName] = useState('');

    // При маунте получаем список всех билетов
    useEffect(() => {
        fetchTickets();
    }, []);

    // Получение списка всех билетов
    const fetchTickets = async () => {
        try {
            const response = await fetch('http://localhost:8080/tickets/all', { method: 'GET' });
            if (!response.ok) {
                console.error('Ошибка при получении списка билетов');
                return;
            }
            const data = await response.json();
            // Предполагается, что вернётся массив объектов
            setTickets(data);
        } catch (error) {
            console.error('Сетевая ошибка при получении билетов:', error);
        }
    };

    // Создание нового билета
    const createTicket = async () => {
        try {
            const requestBody = {
                login: login.trim(),
                event_name: eventName.trim(),
                price: parseFloat(price),
                ticket_type: ticketType,
            };

            const response = await fetch('http://localhost:8080/tickets/create', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(requestBody),
            });

            if (!response.ok) {
                console.error('Ошибка при создании билета');
                return;
            }

            // Обновляем список билетов и очищаем форму
            await fetchTickets();
            clearForm();
        } catch (error) {
            console.error('Сетевая ошибка при создании билета:', error);
        }
    };

    // Получить билет по логину и названию события
    const getTicket = async (userLogin: string, event: string) => {
        try {
            const response = await fetch(`/tickets/${encodeURIComponent(userLogin)}/${encodeURIComponent(event)}`, {
                method: 'GET',
            });
            if (!response.ok) {
                console.error('Ошибка при получении билета');
                return null;
            }
            const ticketData = await response.json();
            return ticketData;
        } catch (error) {
            console.error('Сетевая ошибка при получении билета:', error);
            return null;
        }
    };

    // Начать редактирование билета
    const handleEditClick = async (userLogin: string, event: string) => {
        setIsEditMode(true);
        setOriginalLogin(userLogin);
        setOriginalEventName(event);

        const ticketData = await getTicket(userLogin, event);
        if (ticketData) {
            setLogin(ticketData.login ?? '');
            setEventName(ticketData.event_name ?? '');
            setPrice(String(ticketData.price ?? ''));
            setTicketType(ticketData.ticket_type ?? 0);
        }
    };

    // Обновление данных билета
    const updateTicket = async () => {
        try {
            const requestBody = {
                login: login.trim(),
                event_name: eventName.trim(),
                price: parseFloat(price),
                ticket_type: ticketType,
            };

            // PUT /tickets/{originalLogin}/{originalEventName}
            const response = await fetch(
                `/tickets/${encodeURIComponent(originalLogin)}/${encodeURIComponent(originalEventName)}`,
                {
                    method: 'PUT',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(requestBody),
                }
            );

            if (!response.ok) {
                console.error('Ошибка при обновлении билета');
                return;
            }

            await fetchTickets();
            clearForm();
        } catch (error) {
            console.error('Сетевая ошибка при обновлении билета:', error);
        }
    };

    // Удаление билета
    const deleteTicket = async (userLogin: string, event: string) => {
        try {
            const response = await fetch(
                `/tickets/${encodeURIComponent(userLogin)}/${encodeURIComponent(event)}`,
                {
                    method: 'DELETE',
                }
            );
            if (!response.ok) {
                console.error('Ошибка при удалении билета');
                return;
            }
            await fetchTickets();
        } catch (error) {
            console.error('Сетевая ошибка при удалении билета:', error);
        }
    };

    // Очистка полей формы и сброс режима редактирования
    const clearForm = () => {
        setLogin('');
        setEventName('');
        setPrice('');
        setTicketType(0);
        setIsEditMode(false);
        setOriginalLogin('');
        setOriginalEventName('');
    };

    // Сабмит формы
    const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!login.trim() || !eventName.trim()) {
            alert('Поля "Логин" и "Название события" обязательны!');
            return;
        }

        if (isEditMode) {
            updateTicket();
        } else {
            createTicket();
        }
    };

    return (
        <div className="max-w-3xl mx-auto py-8">
            <h1 className="scroll-m-20 text-2xl font-bold tracking-tight mb-4">Управление билетами</h1>

            {/* Форма для создания/редактирования билета */}
            <form onSubmit={handleFormSubmit} className="space-y-4 mb-8">
                <div>
                    <label className="block mb-1">Логин пользователя:</label>
                    <Input
                        value={login}
                        onChange={(e) => setLogin(e.target.value)}
                        placeholder="Введите логин"
                        required
                    />
                </div>

                <div>
                    <label className="block mb-1">Название события:</label>
                    <Input
                        value={eventName}
                        onChange={(e) => setEventName(e.target.value)}
                        placeholder="Введите название события"
                        required
                    />
                </div>

                <div>
                    <label className="block mb-1">Цена билета (price):</label>
                    <Input
                        type="number"
                        step="0.01"
                        value={price}
                        onChange={(e) => setPrice(e.target.value)}
                        placeholder="Введите цену"
                    />
                </div>

                <div>
                    <label className="block mb-1">Тип билета (ticket_type):</label>
                    <Select value={String(ticketType)} onValueChange={(value) => setTicketType(Number(value))}>
                        <SelectTrigger className="w-full">
                            <SelectValue placeholder="Выберите тип билета" />
                        </SelectTrigger>
                        <SelectContent>
                            <SelectItem value="0">Free (0)</SelectItem>
                            <SelectItem value="1">Paid (1)</SelectItem>
                        </SelectContent>
                    </Select>
                </div>

                <div className="flex items-center gap-2">
                    <Button type="submit" variant="default">
                        {isEditMode ? 'Обновить билет' : 'Создать билет'}
                    </Button>
                    {isEditMode && (
                        <Button variant="ghost" type="button" onClick={clearForm}>
                            Отмена
                        </Button>
                    )}
                </div>
            </form>

            {/* Таблица со списком билетов */}
            <div className="overflow-x-auto">
                <table className="w-full border border-gray-200 rounded-md">
                    <thead className="bg-gray-100">
                    <tr>
                        <th className="px-4 py-2 text-left">Логин</th>
                        <th className="px-4 py-2 text-left">Событие</th>
                        <th className="px-4 py-2 text-left">Цена</th>
                        <th className="px-4 py-2 text-left">Тип билета</th>
                        <th className="px-4 py-2">Действия</th>
                    </tr>
                    </thead>
                    <tbody>
                    {tickets.length > 0 ? (
                        tickets.map((ticket) => (
                            <tr key={`${ticket.login}-${ticket.event_name}`} className="border-b last:border-none">
                                <td className="px-4 py-2">{ticket.login}</td>
                                <td className="px-4 py-2">{ticket.event_name}</td>
                                <td className="px-4 py-2">{ticket.price}</td>
                                <td className="px-4 py-2">
                                    {ticket.ticket_type === 0 ? 'Free (0)' : 'Paid (1)'}
                                </td>
                                <td className="px-4 py-2">
                                    <div className="flex gap-2 justify-center">
                                        <Button
                                            variant="secondary"
                                            onClick={() => handleEditClick(ticket.login, ticket.event_name)}
                                        >
                                            Редактировать
                                        </Button>
                                        <Button
                                            variant="destructive"
                                            onClick={() => deleteTicket(ticket.login, ticket.event_name)}
                                        >
                                            Удалить
                                        </Button>
                                    </div>
                                </td>
                            </tr>
                        ))
                    ) : (
                        <tr>
                            <td colSpan={5} className="px-4 py-6 text-center">
                                Нет билетов
                            </td>
                        </tr>
                    )}
                    </tbody>
                </table>
            </div>
        </div>
    );
}