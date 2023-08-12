from aiogram import Dispatcher, Bot
from aiogram.contrib.fsm_storage.memory import MemoryStorage

storage = MemoryStorage()
bot = Bot(f'TOKEN')
dp = Dispatcher(bot, storage=storage)

texts = ["""Привет, отправь мне фото банки, название энергетика, страну производства и количество кофеина/таурина в нём\nПример:\nBlack Monster Energy Pipeline Punch\nРоссия\nне более 31мг/100мл\n240мг/100мл\nВажно: Присылай состав описанием к фотографии и не используй сжатие\nЕсли возникли проблемы, или нет фотографии, напиши мне: @ekelen""", 'Заявка принята, спасибо']
