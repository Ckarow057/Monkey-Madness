import sqlite3

# Connect to the SQLite database
conn = sqlite3.connect('monkeys.db')
c = conn.cursor()

# Create the monkeys table
c.execute('''CREATE TABLE monkeys
             (monkeyID INTEGER PRIMARY KEY, 
             monkeyName TEXT, 
             monkeyBreed TEXT,
             monkeyImg TEXT, 
             monkeyFact TEXT)''')

c.execute('''CREATE TABLE userinformation
             (userID INTEGER PRIMARY KEY, 
             userSSN INTEGER, 
             userCardNumber INTEGER)''')

# Sample data to populate the table
monkeys_data = [
    (1, 'Japanese Macaque', 'Macaca Fuscata','japanese_macaque.jpg' ,'Japanese Macaques bathe in thermal springs that are heated by nearby volcanoes.'),
    (2, 'Gorilla', 'Gorilla beringei','gorilla.jpg' ,'Gorillas are the largest primates.'),
    (3, 'Orangutan', 'Pongo abelii','orangutan.jpg', 'Orangutans are known for their distinctive red fur.'),
    (4, 'Bonobo', 'Pan paniscus','bonobo.jpg', 'Bonobos are often referred to as pygmy chimpanzees.'),
    (5, "De Brazza's Monkey", 'Cercopithecus neglectus', 'de_brazza_monkey.jpg', "If De Brazza's monkeys sense danger, the group will stand perfectly still while the alpha male shakes tree branches to take attention away from the other group members."),
    (6, "Spider Monkey", 'Ateles', 'spider_monkey.jpg', "Spider Monkeys do not have thumbs."),
    (7, "Red Ruffed Lemur", 'Varecia rubra', 'red_ruffed_lemur.jpg', "All lemurs originate from Madigascar."),
    (8, "Siamang", "Symphalangus Syndactylus", 'siamang.jpg', "Siamangs have an air sack on their necks that allow them to make a whooping noise."),
    (9, "Cotton Top Tamarin", "Saguinus Oedipus", 'cotton_top_tamarin.jpg', "Cotton Top Tamarins only grow to about the size of a squirrel"),
    (10, "Goeldi Monkey", "Callimico Goeldii", 'goeldi_monkey.jpg', "Goeldi monkeys are the only known monkey to regularly eat mushrooms."),
    # Add more data as needed
]

users_info = [
    (1, 391647654, 377456767226730),
    (2, 390372674, 377456036706215),
    (3, 390372234, 377456036723415),
    (4, 390324523, 322356036723415),
    (5, 390321233, 322356036486415)
]

# Insert the data into the table
c.executemany('INSERT INTO monkeys VALUES (?, ?, ?, ?, ?)', monkeys_data)
c.executemany('INSERT INTO userinformation VALUES (?, ?, ?)', users_info)


# Commit changes and close the connection
conn.commit()
conn.close()
