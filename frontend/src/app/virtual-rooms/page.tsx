'use client';

import React, { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

type TVirtualRoom = {
    room_name: string;
    capacity: number;
};

export default function VirtualRoomsPage() {
    // Список виртуальных комнат
    const [rooms, setRooms] = useState<TVirtualRoom[]>([]);

    // Флаг загрузки
    const [loader, setLoader] = useState(true);

    // Поля для формы
    const [roomName, setRoomName] = useState('');
    const [capacity, setCapacity] = useState('');

    // Режим редактирования
    const [isEditMode, setIsEditMode] = useState(false);
    // Исходное название комнаты (чтобы знать, какую комнату мы редактируем)
    const [originalRoomName, setOriginalRoomName] = useState('');

    useEffect(() => {
        fetchRooms();
    }, []);

    // Получение списка всех виртуальных комнат
    const fetchRooms = async () => {
        try {
            const response = await fetch('http://localhost:8080/virtualRooms/all', { method: 'GET' });
            if (!response.ok) {
                console.error('Ошибка при получении списка комнат');
                return;
            }
            const data = await response.json();

            console.log('Response from server:', data); // <-- Посмотрите, что тут

            setRooms(data.rooms);
        } catch (error) {
            console.error('Сетевая ошибка при получении комнат:', error);
        } finally {
            setLoader(false);
        }
    };

    // Создание новой комнаты
    const createRoom = async () => {
        try {
            const requestBody = {
                room_name: roomName.trim(),
                capacity: parseInt(capacity, 10),
            };

            const response = await fetch('http://localhost:8080/virtualRooms/create', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(requestBody),
            });

            if (!response.ok) {
                console.error('Ошибка при создании комнаты');
                return;
            }
            // После создания комнаты заново получаем список
            setLoader(true);
            await fetchRooms();
            clearForm();
        } catch (error) {
            console.error('Сетевая ошибка при создании комнаты:', error);
        }
    };

    // Получить данные комнаты по названию
    const getRoomByName = async (name: string) => {
        try {
            const response = await fetch(`http://localhost:8080/virtualRooms/${encodeURIComponent(name)}`, { method: 'GET' });
            if (!response.ok) {
                console.error('Ошибка при получении комнаты');
                return null;
            }
            return await response.json();
        } catch (error) {
            console.error('Сетевая ошибка при получении комнаты:', error);
            return null;
        }
    };

    // Начало процесса редактирования
    const handleEditClick = async (name: string) => {
        setIsEditMode(true);
        setOriginalRoomName(name);

        const roomData = await getRoomByName(name);
        if (roomData) {
            setRoomName(roomData.room_name ?? '');
            setCapacity(String(roomData.capacity ?? ''));
        }
    };

    // Обновление комнаты
    const updateRoom = async () => {
        try {
            const requestBody = {
                room_name: roomName.trim(),
                capacity: parseInt(capacity, 10),
            };

            const response = await fetch(`http://localhost:8080/virtualRooms/update/${encodeURIComponent(originalRoomName)}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(requestBody),
            });

            if (!response.ok) {
                console.error('Ошибка при обновлении комнаты');
                return;
            }
            // После обновления заново получаем список
            setLoader(true);
            await fetchRooms();
            clearForm();
        } catch (error) {
            console.error('Сетевая ошибка при обновлении комнаты:', error);
        }
    };

    // Удаление комнаты
    const deleteRoom = async (name: string) => {
        try {
            const response = await fetch(`/virtualRooms/${encodeURIComponent(name)}`, {
                method: 'DELETE',
            });
            if (!response.ok) {
                console.error('Ошибка при удалении комнаты');
                return;
            }
            // После удаления снова загружаем список
            setLoader(true);
            await fetchRooms();
        } catch (error) {
            console.error('Сетевая ошибка при удалении комнаты:', error);
        }
    };

    // Очистка формы и сброс режима
    const clearForm = () => {
        setRoomName('');
        setCapacity('');
        setIsEditMode(false);
        setOriginalRoomName('');
    };

    // Сабмит формы
    const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!roomName.trim()) {
            alert('Поле "Название комнаты" обязательно');
            return;
        }

        if (isEditMode) {
            updateRoom();
        } else {
            createRoom();
        }
    };

    // Если данные ещё в процессе загрузки — показываем лоадер
    if (loader) {
        return (
            <div className="max-w-3xl mx-auto py-8">
                <h1 className="scroll-m-20 text-2xl font-bold tracking-tight mb-4">Управление виртуальными комнатами</h1>
                <div className="text-center text-gray-500">Загрузка...</div>
            </div>
        );
    }

    return (
        <div className="max-w-3xl mx-auto py-8">
            <h1 className="scroll-m-20 text-2xl font-bold tracking-tight mb-4">Управление виртуальными комнатами</h1>

            {/* Форма создания / редактирования комнаты */}
            <form onSubmit={handleFormSubmit} className="space-y-4 mb-8">
                <div>
                    <label className="block mb-1">Название комнаты:</label>
                    <Input
                        value={roomName}
                        onChange={(e) => setRoomName(e.target.value)}
                        placeholder="Введите название комнаты"
                        required
                    />
                </div>

                <div>
                    <label className="block mb-1">Вместимость (capacity):</label>
                    <Input
                        type="number"
                        value={capacity}
                        onChange={(e) => setCapacity(e.target.value)}
                        placeholder="Введите число (вместимость)"
                    />
                </div>

                <div className="flex items-center gap-2">
                    <Button type="submit" variant="default">
                        {isEditMode ? 'Обновить комнату' : 'Создать комнату'}
                    </Button>
                    {isEditMode && (
                        <Button variant="ghost" type="button" onClick={clearForm}>
                            Отмена
                        </Button>
                    )}
                </div>
            </form>

            {/* Таблица списков комнат */}
            <div className="overflow-x-auto">
                <table className="w-full border border-gray-200 rounded-md">
                    <thead className="bg-gray-100">
                    <tr>
                        <th className="px-4 py-2 text-left">Название</th>
                        <th className="px-4 py-2 text-left">Вместимость</th>
                        <th className="px-4 py-2">Действия</th>
                    </tr>
                    </thead>
                    <tbody>
                    {rooms.length > 0 ? (
                        rooms.map((room) => (
                            <tr key={room.room_name} className="border-b last:border-none">
                                <td className="px-4 py-2">{room.room_name}</td>
                                <td className="px-4 py-2">{room.capacity}</td>
                                <td className="px-4 py-2">
                                    <div className="flex gap-2 justify-center">
                                        <Button variant="secondary" onClick={() => handleEditClick(room.room_name)}>
                                            Редактировать
                                        </Button>
                                        <Button variant="destructive" onClick={() => deleteRoom(room.room_name)}>
                                            Удалить
                                        </Button>
                                    </div>
                                </td>
                            </tr>
                        ))
                    ) : (
                        <tr>
                            <td colSpan={3} className="px-4 py-6 text-center">
                                Нет виртуальных комнат
                            </td>
                        </tr>
                    )}
                    </tbody>
                </table>
            </div>
        </div>
    );
}