-- Seed: Users
INSERT INTO user_service.users (id, email, firstname, lastname, firebase_uid)
VALUES ('71b55371-da0e-438f-a689-40ae99c888ce', 'marcela.drnkova@poplk.com', 'Marcela', 'Drnková',
        'kSq0h8K2GrUtiZcegdezuUo7cnq2'),
       ('29bfae54-57ed-422d-8d09-479d01ee7503', 'karel.vrzal@poplk.com', 'Karel', 'Vrzal',
        'bv8WQ5gKlJOJaQQ7KcTnya6lbQA3'),
       ('bae2bf16-8b42-4554-8edd-191e42fcebc4', 'ema.novotna@poplk.com', 'Ema', 'Novotná',
        'mWGaJglM3GTAbNtVL1LLfYPHNpA2'),
       ('504484f9-67c2-4112-b84d-489ee9176b60', 'vojtech.krtek@poplk.com', 'Vojtěch', 'Krtek',
        'ec5TY222U4P83FZtZxlWeP2ZBEz1'),
       ('a26c202a-2d5e-47d9-9728-688d788cb72f', 'adelka.rybickova@poplk.com', 'Adélka', 'Rybičková',
        'm4E2PWU2CjNLOhLmw077hPjHNob2');


-- Seed: Newsletters
INSERT INTO newsletter_service.newsletters (id, title, description, created_at, user_id)
VALUES ('a1c1e8e0-1e2f-4b6b-8b8f-aaa111111aaa', 'Ranní káva s Marcelou',
        'Krátké ranní zamyšlení nad každodenním chaosem.', now(), '71b55371-da0e-438f-a689-40ae99c888ce'),
       ('b2d2e9f1-2f3a-4c7c-9c9a-bbb222222bbb', 'Kód a klávesy',
        'Karel sdílí zajímavosti ze světa backendu a syntaktických pekel.', now(),
        '29bfae54-57ed-422d-8d09-479d01ee7503'),
       ('c3f3a0a2-3a4b-4d8d-ada3-ccc333333ccc', 'Design v pohybu', 'Ema přináší pohled na UX a UI s šálkem stylu.',
        now(), 'bae2bf16-8b42-4554-8edd-191e42fcebc4'),
       ('d4a4b1b3-4b5c-4e9e-bee4-ddd444444ddd', 'Z lesní nory', 'Vojtěch píše o přírodě, klidu a digitálním detoxu.',
        now(), '504484f9-67c2-4112-b84d-489ee9176b60'),
       ('e5b5c2c4-5c6d-5f0f-cff5-eee555555eee', 'Drobnosti s Adélkou',
        'Adélka sdílí malé radosti a kreativní tipy na každý týden.', now(), 'a26c202a-2d5e-47d9-9728-688d788cb72f');

-- Seed: Posts
INSERT INTO newsletter_service.posts (id, newsletter_id, title, content, html_content, published, created_at)
VALUES
-- Marcela
('f6d6e3e5-6d7e-6101-d001-fff666666fff', 'a1c1e8e0-1e2f-4b6b-8b8f-aaa111111aaa', 'Pondělní restart',
 'Dnes začínáme s hlubokým nádechem a kávou navíc.', '<p>Dnes začínáme s hlubokým nádechem a kávou navíc.</p>', false,
 now()),

-- Karel
('b9d2e8f5-33cf-4e93-8df4-06dc9d75b12f', 'b2d2e9f1-2f3a-4c7c-9c9a-bbb222222bbb', 'Když se API vzbouří',
 'Minulý týden jsme lovili bugy v produkci – detaily uvnitř.',
 '<p>Minulý týden jsme lovili bugy v produkci – detaily uvnitř.</p>', false, now()),

-- Ema
('3f6f1ab0-29a2-4bc3-bf2c-3768d26c42f7', 'c3f3a0a2-3a4b-4d8d-ada3-ccc333333ccc', 'Barvy, které ladí',
 'Jak kombinovat barvy, aby z toho nešly oči kolem?', '<p>Jak kombinovat barvy, aby z toho nešly oči kolem?</p>', false,
 now()),
('8c7a9879-70ae-4d6b-a06b-d6c1850a8cb9', 'c3f3a0a2-3a4b-4d8d-ada3-ccc333333ccc', 'UX pro lidi, ne roboty',
 'Příběh o tom, proč přístupnost není volitelná.', '<p>Příběh o tom, proč přístupnost není volitelná.</p>', false,
 now()),

-- Vojtěch
('17d9e7a1-05d0-407a-a889-5cb2f72e5e0f', 'd4a4b1b3-4b5c-4e9e-bee4-ddd444444ddd', 'Šum lesa do inboxu',
 'Jak příroda léčí i bez signálu.', '<p>Jak příroda léčí i bez signálu.</p>', false, now()),

-- Adélka
('a59dd1e3-9e86-4f3e-b7ce-b91f3f8dc11b', 'e5b5c2c4-5c6d-5f0f-cff5-eee555555eee', 'Papírové radosti',
 'DIY nápady na hezké víkendové tvoření.', '<p>DIY nápady na hezké víkendové tvoření.</p>', false, now());
