from aiogram.types import Message
from aiogram import types, Dispatcher
from config import bot, dp, texts


async def start_message(message: Message):
    await bot.send_message(message.chat.id, texts[0])


async def send_photo(message):
    await bot.send_photo(f'{group_id}', message.photo[-1].file_id, caption=message.caption)
    await bot.send_message(f'{group_id}', f'@{message.from_user.username}')
    await bot.send_message(message.chat.id, texts[1])

def init_main_chat(dp: Dispatcher):
    dp.register_message_handler(start_message, commands=['start']) 
    dp.register_message_handler(send_photo, content_types=['photo'])
