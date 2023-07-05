import psycopg2

class BasicCommands:
    def __init__(self, dbname, username, password, connection=None, cursor = None):
        self.dbname = dbname
        self.username = username
        self.password = password
        self.connection = connection
        self.cursor = cursor

    def connect(self):
        self.connection = psycopg2.connect(f"dbname={self.dbname} user={self.username} password={self.password}")
        self.cursor = self.connection.cursor()

    def select_query(self, table):
        self.cursor.execute(f"SELECT * FROM {table}")

    def fetch_all(self):
        fetch = self.cursor.fetchall()
        print(fetch)

    def insert_into(self, table, *args):
        try:
            insert = f"""INSERT INTO {table} VALUES {args}"""
            self.cursor.execute(insert)
            self.connection.commit()
        except Exception as e:
            print("[ERROR] ",e)

    def close_connection(self):
        self.cursor.close()
