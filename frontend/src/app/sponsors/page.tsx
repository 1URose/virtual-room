'use client';

import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';   // компоненты shadcn/ui
import { Input } from '@/components/ui/input';

type TSponsor = {
    name: string;
    contact_info: string;
    contribution_amount: number;
    event_name: string;
};

// Наш клиентский компонент
export default function SponsorsPage() {
    const [sponsors, setSponsors] = useState<TSponsor[]>([]);

    // Поля формы
    const [name, setName] = useState('');
    const [contactInfo, setContactInfo] = useState('');
    const [contributionAmount, setContributionAmount] = useState('');
    const [eventName, setEventName] = useState('');

    // Режим редактирования
    const [isEditMode, setIsEditMode] = useState(false);
    // Старое имя спонсора (для update/delete)
    const [originalName, setOriginalName] = useState('');

    // При маунте компонента загружаем список спонсоров
    useEffect(() => {
        fetchSponsors();
    }, []);

    // Запрос списка спонсоров
    const fetchSponsors = async () => {
        try {
            const response = await fetch('http://localhost:8080/sponsors/all', {
                method: 'GET',
            });
            if (!response.ok) {
                console.error('Ошибка при получении списка спонсоров');
                return;
            }
            const data = await response.json();
            // Предполагается, что вернётся массив объектов
            setSponsors(data);
        } catch (error) {
            console.error('Сетевая ошибка при получении спонсоров:', error);
        }
    };

    // Создание спонсора
    const createSponsor = async () => {
        try {
            const requestBody = {
                name: name.trim(),
                contact_info: contactInfo.trim(),
                contribution_amount: parseFloat(contributionAmount),
                event_name: eventName.trim(),
            };

            const response = await fetch('http://localhost:8080/sponsors/create', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(requestBody),
            });

            if (!response.ok) {
                console.error('Ошибка при создании спонсора');
                return;
            }

            await fetchSponsors();
            clearForm();
        } catch (error) {
            console.error('Сетевая ошибка при создании спонсора:', error);
        }
    };

    // Получить спонсора по имени
    const getSponsorByName = async (sponsorName: string) => {
        try {
            const response = await fetch(`/sponsors/${encodeURIComponent(sponsorName)}`, {
                method: 'GET',
            });
            if (!response.ok) {
                console.error('Ошибка при получении спонсора');
                return null;
            }
            return await response.json();
        } catch (error) {
            console.error('Сетевая ошибка при получении спонсора:', error);
            return null;
        }
    };

    // Начать редактирование (загружаем данные, подставляем в форму)
    const handleEditClick = async (sponsorName: string) => {
        setIsEditMode(true);
        setOriginalName(sponsorName);
        const sponsorData = await getSponsorByName(sponsorName);
        if (sponsorData) {
            setName(sponsorData.name ?? '');
            setContactInfo(sponsorData.contact_info ?? '');
            setContributionAmount(String(sponsorData.contribution_amount ?? ''));
            setEventName(sponsorData.event_name ?? '');
        }
    };

    // Обновление спонсора
    const updateSponsor = async () => {
        try {
            const requestBody = {
                name: name.trim(),
                contact_info: contactInfo.trim(),
                contribution_amount: parseFloat(contributionAmount),
                event_name: eventName.trim(),
            };

            const response = await fetch(`/sponsors/${encodeURIComponent(originalName)}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(requestBody),
            });

            if (!response.ok) {
                console.error('Ошибка при обновлении спонсора');
                return;
            }
            await fetchSponsors();
            clearForm();
        } catch (error) {
            console.error('Сетевая ошибка при обновлении спонсора:', error);
        }
    };

    // Удаление спонсора
    const deleteSponsor = async (sponsorName: string) => {
        try {
            const response = await fetch(`/sponsors/${encodeURIComponent(sponsorName)}`, {
                method: 'DELETE',
            });
            if (!response.ok) {
                console.error('Ошибка при удалении спонсора');
                return;
            }
            await fetchSponsors();
        } catch (error) {
            console.error('Сетевая ошибка при удалении спонсора:', error);
        }
    };

    // Очистка формы
    const clearForm = () => {
        setName('');
        setContactInfo('');
        setContributionAmount('');
        setEventName('');
        setIsEditMode(false);
        setOriginalName('');
    };

    // Обработчик формы
    const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!name.trim()) {
            alert('Поле "Имя спонсора" обязательно');
            return;
        }
        if (isEditMode) {
            updateSponsor();
        } else {
            createSponsor();
        }
    };

    return (
        <div className="max-w-3xl mx-auto py-8">
            <h1 className="scroll-m-20 text-2xl font-bold tracking-tight mb-4">Управление спонсорами</h1>

            {/* Форма создания/редактирования */}
            <form onSubmit={handleFormSubmit} className="space-y-4 mb-8">
                <div>
                    <label className="block mb-1">Имя спонсора:</label>
                    <Input
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        placeholder="Введите имя спонсора"
                        required
                    />
                </div>

                <div>
                    <label className="block mb-1">Контактная информация:</label>
                    <Input
                        value={contactInfo}
                        onChange={(e) => setContactInfo(e.target.value)}
                        placeholder="Введите контактную информацию"
                    />
                </div>

                <div>
                    <label className="block mb-1">Сумма вклада:</label>
                    <Input
                        type="number"
                        step="0.01"
                        value={contributionAmount}
                        onChange={(e) => setContributionAmount(e.target.value)}
                        placeholder="Введите сумму"
                    />
                </div>

                <div>
                    <label className="block mb-1">Имя события:</label>
                    <Input
                        value={eventName}
                        onChange={(e) => setEventName(e.target.value)}
                        placeholder="К какому событию относится спонсор"
                    />
                </div>

                <div className="flex items-center gap-2">
                    <Button type="submit" variant="default">
                        {isEditMode ? 'Обновить спонсора' : 'Создать спонсора'}
                    </Button>
                    {isEditMode && (
                        <Button variant="ghost" type="button" onClick={clearForm}>
                            Отмена
                        </Button>
                    )}
                </div>
            </form>

            {/* Таблица со всеми спонсорами */}
            <div className="overflow-x-auto">
                <table className="w-full border border-gray-200 rounded-md">
                    <thead className="bg-gray-100">
                    <tr>
                        <th className="px-4 py-2 text-left">Имя</th>
                        <th className="px-4 py-2 text-left">Контакты</th>
                        <th className="px-4 py-2 text-left">Вклад</th>
                        <th className="px-4 py-2 text-left">Событие</th>
                        <th className="px-4 py-2">Действия</th>
                    </tr>
                    </thead>
                    <tbody>
                    {sponsors.length > 0 ? (
                        sponsors.map((sponsor) => (
                            <tr key={sponsor.name} className="border-b last:border-none">
                                <td className="px-4 py-2">{sponsor.name}</td>
                                <td className="px-4 py-2">{sponsor.contact_info}</td>
                                <td className="px-4 py-2">{sponsor.contribution_amount}</td>
                                <td className="px-4 py-2">{sponsor.event_name}</td>
                                <td className="px-4 py-2">
                                    <div className="flex gap-2 justify-center">
                                        <Button variant="secondary" onClick={() => handleEditClick(sponsor.name)}>
                                            Редактировать
                                        </Button>
                                        <Button variant="destructive" onClick={() => deleteSponsor(sponsor.name)}>
                                            Удалить
                                        </Button>
                                    </div>
                                </td>
                            </tr>
                        ))
                    ) : (
                        <tr>
                            <td colSpan={5} className="px-4 py-6 text-center">
                                Нет спонсоров
                            </td>
                        </tr>
                    )}
                    </tbody>
                </table>
            </div>
        </div>
    );
}