from aiogram import executor
from config import dp
from main import init_main_chat


async def on_startup(_):
    print('Я работаю, хозяинка')



init_main_chat(dp)
executor.start_polling(dp, skip_updates=True)
