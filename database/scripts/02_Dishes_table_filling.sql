ALTER SEQUENCE dishes_dish_id_seq RESTART WITH 1;

insert into dishes
values
-- https://calorizator.ru/recipe/18351
(0, 'Фруктовый салат с йогуртом',
 'Шаг 1: Киви почистить и порезать кубиками. Шаг 2: Мандарины очистить, разделить на дольки, убрать пленку и порезать кубиками. Шаг 3: Грушу очистить от кожуры, убрать сердцевину, порезать кубиками. Шаг 4: Банан очистить и нарезать кубиками. Шаг 5: Съесть сразу, в холодильник не ставить и не хранить.',
 3.5, 23.5, 1.6, 119.9, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipe/18351.jpg', '0ec197ebfcbcf8768574c4faf15d197fac861c323981f46a80236274efd7d4b7');

insert into dishes (dish_name, recipe, proteins, carbs, fats, calories, image_path, dish_hash)
values
-- https://calorizator.ru/recipes/182108
('Кокосовые сырники',
 'Шаг 1: В глубокую миску выложить творог, сахар, кокосовую стружку и яйца. Все смешать. Шаг 2: Добавить муку. Шаг 3: Тщательно размять смесь руками, она должна хорошо держать форму. Шаг 4: Слепить 12 небольших сырников. Шаг 5: Обжарить сырники на сухой сковороде с обеих сторон по 4-5 минут.',
 24.7, 18.7, 24.3, 394.4, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/182108.jpg', '7b63102c93aacb68ec46e9b0e8913f9f9f0c7496ca9b293a4a45abbe2751f220'),

-- https://calorizator.ru/recipes/182106
('Закуска из брокколи с сыром',
 'Шаг 1: Сыр натереть на терке. Шаг 2: Брокколи отварить до полуготовности. Шаг 3: Выложить соцветия на пергамент, приплюснуть стаканом. Шаг 4: Выложить на каждое соцветие маленький кусочек сливочного масла. Шаг 5: Присыпать перцем и сухим чесноком. Шаг 6: Сверху выложить сыр. Шаг 7: Запекать в духовке 7 минут с конвекцией.',
 5.7, 5.3, 9.9, 126.6, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/182106.jpg', 'a42482a425bffb82f55585fbf4b24cd8ac3047366648e44a77ce8501f01dc5ec'),

-- https://calorizator.ru/recipes/170442
('Оладушки на мягком твороге',
 'Шаг 1: Яйца взбить в пену. Шаг 2: Добавить мягкий творог и тщательно перемешать. Шаг 3: Добавить муку и разрыхлитель и перемешать, дать постоять 10 минут. Шаг 4: Жарить на среднем огне до готовности.',
 10.3, 21, 2.4, 150.2, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/170442.jpg', '7a13938e568e4d32f80bc4b091b8965fae8f750add327b4e2ed8efcd8fa9c086'),

-- https://calorizator.ru/recipes/114930
('Ленивая овсянка с черникой и протеином',
 'Шаг 1: Промыть овсянку под проточной водой для улучшения вкуса каши. Шаг 2: Засыпать её в баночку. Шаг 3: Взбить протеин и молоко блендером. Шаг 4: Добавить молочную смесь, йогурт и ягоды к овсянке, всё хорошо перемешать. Шаг 5: Поставить на ночь в холодильник для набухания овсянки. Шаг 6: Утром перемешать и съесть холодной или подогреть, украсить ягодами.',
 26, 38.1, 7.5, 325.4, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/114930.jpg', 'bba715e1923cabbb7312415583fd905ad480bcaacb004c1b015fda4c9538b8f7'),

-- https://calorizator.ru/recipes/117765
('Домашние протеиновые батончики',
 'Шаг 1: Орехи и цукаты измельчить в крошку. Шаг 2: Протеин смешать с сухим молоком. Шаг 3: Добавить ореховую смесь и постепенно добавляя воду замесить пластичное тесто. Шаг 4: Слепить батончики.',
 27, 38.3, 4.2, 303.4, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/117765.jpg', '79b9b00f8d4e288668c42c1207cdb9fe4aed83285c082f4129dcc3b105390a88'),

-- https://calorizator.ru/recipes/113034
('Творожный десерт с бананом ПП',
 'Шаг 1: Овсянку смолоть в муку блендером. Шаг 2: Растереть овсяную муку с творогом. Шаг 3: Вскипятить молоко, добавить какао и сахарозаменитель, хорошо перемешать до полного растворения какао. Шаг 4: Добавить молоко к творогу, размешать. Шаг 5: Выложить творожную массу на пищевую пленку, разровнять, уложить в центр банан. Шаг 6: Свернуть массу так, чтобы банан был внутри, убрать в морозилку на 20 минут.',
 11.1, 14.5, 1.6, 114.3, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/113034.jpg', '1d89bf87a4f81822ddd65ae9abf8a16da50164a070c16363bdcb68b04faf19b0'),


-- https://calorizator.ru/recipes/131893
('Шарлотка ПП без муки и сахара',
 'Шаг 1: Овсяные хлопья смолоть в муку. Шаг 2: Яйца взбить с медом. Шаг 3: Соединить сухие и влажные ингредиенты. Шаг 4: Яблоки очистить от сердцевины. Шаг 5: Нарезать яблоки кубиками. Шаг 6: Смешать тесто с яблоками. Шаг 7: Выложить в форму и поставить в духовку на 40 минут при 180 градусах.',
 6.7, 26.9, 5.2, 179.6, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/131893.jpg', '6f891420dec08c317b4c6901caafc0fa0beec498b0eab61cd7332b3261c7da71'),

-- https://calorizator.ru/recipes/140142
('Десерт из овсянки с яблоком',
 'Шаг 1: Овсяные хлопья обжарить на сухой сковороде до золотистого цвета и орехового аромата. Шаг 2: Яблоко вымыть, очистить от кожуры и сердцевины и нарезать небольшими дольками. Шаг 3: Кусочки яблока посыпать сверху сахаром и корицей. Шаг 4: Запечь в духовке при температуре 200 °С 10-15 мин. или в микроволновой печи при 800 Вт 3 мин. до мягкости. Шаг 5: Финики помыть, убрать косточки и порезать на кусочки. Шаг 6: Смешать овсяные хлопья, молоко, яблоки и финики. Шаг 7: Поставить в холод на ночь. Шаг 8: Взбить блендером до однородного состояния.',
 2.7, 21.6, 2.3, 114.1, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/140142.jpg', 'd21712419cc489d7fa73daa2ce745875c363352168c56b11dfa74533a463833e'),

-- https://calorizator.ru/recipes/118827
('Творожно-яблочное суфле ПП',
 'Шаг 1: Яблоко и грушу нарезать небольшими кубиками. Шаг 2: Довести до кипения на антипригарной сковороде. Шаг 3: Творог взбить с йогуртом и сахарозаменителем. Шаг 4: Выложить творожную массу в антипригарную форму, разровнять. Шаг 5: Белки взбить с сахарозаменителем до крутых пик. Шаг 6: Поверх творожного слоя в форму уложить остуженные фрукты. Шаг 7: Сверху выложить взбитые белки. Шаг 8: Выпекать в духовке 30-35 минут при температуре 160 градусов.',
 7, 6.4, 0.4, 56, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/118827.jpg', '997f070b22d4cb8dcd91ece4a8ce1ecbe7787672d3ec510019252c9744b82b00'),

-- https://calorizator.ru/recipes/89884
( 'А-ля творожные сырки из пшенки',
 'Шаг 1: Отварить пшено до готовности. Шаг 2: Залить изюм на 20 минут. Шаг 3: Натереть цедру. Шаг 4: Выжать сок из лайма. Шаг 5: Смешать пшено с соком лайма, цедрой, сиропом агавы. Шаг 6: Растопить кокосовую пасту (и она в твердом состоянии). Шаг 7: Добавить кокосовую стружку и пасту с изюмом в массу и тщательно перемешать. Шаг 8: Разложить по формочкам и отправить в морозилку застыть. Шаг 9: В это время растопить тертое какао на водяной бане. Шаг 10: Достать сырки и глазировать.',
 5.9, 28.4, 8.9, 212.9, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/89884.jpg', '72e75488856681f198c348b11534f76808ba0fb62f14fd4f9564518742f9738a'),

-- https://calorizator.ru/recipes/59057
( 'Айва с орехами в духовке',
 'Шаг 1: Айву залить водой и хорошо промыть. Избавится от ворсинок с помощью щетки или просто хорошо промыть руками. Шаг 2: Верхушку айвы срезать, а сердцевину вырезать ножиком. Избавится от косточек, не повредив шкурку айвы. Шаг 3: Орехи измельчить в блендере. К ним добавить корицу и сахар, все хорошо перемешать. Шаг 4: Айву выложить в форму, застеленную пергаментной бумагой, и заполнить ореховой смесью. Шаг 5: Сверху положить кусочки масла, чтоб айва была сочнее. Шаг 6: Запекать в духовке 50 минут при 180 градусах. Готовую айву нужно остудить и посыпать сахарной пудрой. Шаг 7: При желании украсить веточками мяты.',
 2.2, 35.6, 14.8, 277.3, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/59057.jpg', '1dd2bb40436b3adb49116b2590d254c82978a451b4e017d7735f9427d1ff4eae'),

-- https://calorizator.ru/recipes/66233
( 'Апельсиново-кукурузный пирог',
 'Шаг 1: Взбить в пышный крем размягченное сливочное масло и сахар. Шаг 2: Ввести одно за другим яйца, добавляя с каждым по 1 ст.л. пшеничной муки, чтобы масса не свернулась. Шаг 3: Просеять оставшуюся пшеничную и кукурузную муку. Всыпать смесь вместе с разрыхлителем в тесто. Шаг 4: Тщательно вымыть апельсин и лимон. Натереть на мелкой терке цедру и отжать сок, добавить натертую цедру в тесто. Шаг 5: Отлить 100 мл смешанного сока апельсина и лимона. Остаток добавить в тесто и хорошо размешать. Шаг 6: Влить тесто в глубокую круглую силиконовую форму диаметром 18 см. Разровнять поверхность, поставить форму на противень. Разогреть духовку до 180 °С. Шаг 7: Для глазури смешать в небольшой кастрюльке сахар, 100 мл сока апельсина и лимона и мелко нарезанный мармелад. Шаг 8: На слабом огне довести до кипения и варить около 5 минут, пока глазурь не станет сиропообразной. Готовую глазурь процедить. Шаг 9: Выпекать пирог в разогретой духовке около 1 часа. Проверить готовность пирога с помощью деревянной спицы: после прокалывания центра изделия спица должна быть сухой. Шаг 10: Наколоть готовый пирог спицей по всей поверхности и полить глазурью. Дать изделию остыть, не вынимая из формы. Подавать с мороженым или взбитыми сливками.',
 3.3, 40.1, 12.8, 288.7, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/66233.jpg', 'e36651debe38739046a735016f57a7c7f4cc7ed0d2c48d11e8ef4acd0cd7accd'),

-- https://calorizator.ru/recipes/61590
( 'Арахисовые сырники в духовке',
 'Шаг 1: Замочить изюм в горячей воде. Шаг 2: Смолоть отруби клетчатку в кофемолке. Шаг 3: Смешать все ингредиенты. Шаг 4: Скатать шарики и немного приплюснуть. Шаг 5: Выложить на противень и выпекать в духовке 25 минут при 180 градусах.',
 15.2, 18, 6.7, 192.2, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/61590.jpg', 'cdcb1c86616de696ed289c77d4c21c8800be1ace41b08d2b2b3c97539764bbc7'),

-- https://calorizator.ru/recipes/136099
( 'Белковые панкейки',
 'Шаг 1: Отварить тыкву в течение 13 минут. Шаг 2: Смешать три вида муки.Шаг 3: Белки отделить от желтков. Шаг 4: Белки взбить в однородную массу с тыквой. Шаг 5: Добавить в пюре муку и тщательно перемешать. Шаг 6: Обжарить панкейки с двух сторон (т.л. – 1 панкейк).',
 9.4, 13.9, 3.1, 119.5, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/136099.jpg', 'e958655c553122affb3d8e448a8acea4c2ca61728b7a4e7cef1511f3505bca57'),

-- https://calorizator.ru/recipes/37840
( 'Бисквитное печенье',
 'Шаг 1: Охлажденные яйца взбить с сахаром и ванилью в очень плотную пену. Взбивать на высокой скорости миксера. Все кристаллики сахара должны полностью растереться с яйцами. Шаг 2: Постепенно в яично-сахарную пену ввести муку, при этом скорость миксера уменьшить до минимума. Шаг 3: Должно получиться густое тесто, консистенция как у густой жирной сметаны. Тесто похоже на обычное бисквитное тесто. Шаг 4: Поместить тесто в кондитерский мешок и отсадить лепешки. Совершенно не важно, какую насадку вы выберете, тесто все равно расплывется в лепешку. Шаг 5: Выпекать печенье при температуре 180 градусов 10-15 минут, до появления румяности по краям печенья. Лист пергамента снять с противня и дать печенью на нем полностью остыть, после чего снять при помощи лопатки или широкого ножа.',
 8.8, 58.2, 3.7, 299.6, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/37840.jpg', 'c96ea5967c1504ec8a1e1974120cd82caa7b3fc2d43f7e304c16e01afe5f5c79'),

-- https://calorizator.ru/recipes/59235
( 'Блинный домик с джемом и орехами',
 'Шаг 1: Взбить яйца с солью и сахаром. Шаг 2: Влить 100 мл молока. Шаг 3: Добавить просеянную муку и перемешать с помощью миксера. Шаг 4: Затем постепенно разбавить оставшимся молоком. Влить 2 ст.л. растительного масла, размешать. Шаг 5: Разогреть сковороду и смазать маслом. Из получившегося теста испечь небольшие блины. Шаг 6: Орехи измельчить в чаше блендера или порубить ножом. Шаг 7: Соединить орехи с джемом и перемешать. Шаг 8: Промазать начинкой каждый блин и свернуть трубочкой. Шаг 9: Сложить блины с начинкой друг на друга, чтобы получилась пирамидка. Шаг 10: Взбить миксером сметану с сахаром, полить сметанным кремом блины и убрать в холодильник на 1 час. Перед подачей украсить десерт тертым шоколадом.',
 4.8, 33.7, 13.7, 275.5, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/59235.jpg', '02208adce8404d94eea6245a7c89cf705d4b25fa4281d748723572452626be03'),

-- https://calorizator.ru/recipes/55865
( 'Брускетта с сыром',
 'Шаг 1: Чеснок очистить, пропустить через пресс. Шаг 2: Смешать оливковое масло и чеснок. Шаг 3: Помидор разрезать пополам и нарезать тонкими ломтиками. Шаг 4: Сыр нарезать порционными кусочками. Шаг 5: Зелень кинзы измельчить. Шаг 6: Кусочки багета намазать смесью оливкового масла с чесноком (вух сторон). Шаг 7: Обжарить на сковороде, перевернуть. Шаг 8: Разложить сверху ломтики помидора. Шаг 9: На помидор выложить сыр. На среднем огне под крышкой поставить до расплавления сыра. Шаг 10: При подаче посыпать кинзой.',
 7.7, 16.6, 16, 252, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/55865.jpg', 'b1382b0275e25497745a4972fd4cdf0f7dba567308f5b98525a239773b50c058'),

-- https://calorizator.ru/recipes/52792
( 'Ватрушки с творогом',
 'Шаг 1: В тёплое молоко добавить дрожжи, сахар и соль. Всё тщательно перемешать и оставить на 10 минут до появления шапочки или пузырей. Шаг 2: Влить растительное масло и взбитое яйцо. Размешать. Шаг 3: Просеять муку, но не всю сразу, а постепенно. Шаг 4: Замесить мягкое тесто, которое не липнет к рукам. Оставить тесто подходить в тёплом месте на 1,5-2 часа. Шаг 5: После 1,5 часов тесто должно хорошо подняться. Шаг 6: Тесто разделить на 11 равных шариков примерно по 60 грамм каждый. Шаг 7: Разложить ватрушки на противень посыпанный мукой, в центре каждой ватрушки сделать небольшое углубление для начинки и оставить подходить еще на 30-40 минут в тёплом месте. Шаг 8: Начинка: смешать творог с сахаром, ванилином и куриным желтком. Вымешать массу до однородности. Шаг 9: Выложить творожную начинку в каждую ватрушку, тесто по краям смазать взбитым желтком. Выпекать в заранее разогретой духовке до 180 градусов минут 40 до появления золотистого цвета.',
 9.9, 35.6, 6.9, 244, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/52792.jpg', '6c70b3d103f68c5d802f789c8ccbefca44a8e6d88442066e588559ac6c23a78f'),

-- https://calorizator.ru/recipes/60770
( 'Грибная запеканка',
 'Шаг 1: Грибы помыть и почистить. Шаг 2: Залить грибы горячей водой, посолить и проварить 5 минут. Шаг 3: Выложить грибы шумовкой на полотенце, дать остыть, и порезать на небольшие кусочки. Шаг 4: Луковицу очистить и порезать полукольцами. Шаг 5: Обжарить лук в растительном масле, добавить грибы и слегка обжарить. Шаг 6: Яйца взбить вилкой. Шаг 7: Добавить к яйцам укроп и приправы, перемешать. Шаг 8: Грибы выложить на дно формы смазанной маслом. Шаг 9: Залить яйцами, добавить лавровый лист. Шаг 10: Выпекать при температуре 200 градусов 25 минут.',
 5.6, 1.9, 10.3, 121.8, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/60770.jpg', '91fd18b5ff267c7bd181dad3c54e9d4545ee094e18f66df2c055d2567298f650'),

-- https://calorizator.ru/recipes/85562
( 'Булгур со свининой и перцем',
 'Шаг 1: Булгур обжарить на растительном масле несколько минут. Шаг 2: Овощи нарезать: лук – кубиками, перец – соломкой, морковь – полукольцами. Шаг 3: Булгур залить водой, посолить и поперчить. Шаг 4: Мясо, нарезанное длинными кусочками, обжарить на сковороде, добавить овощи, чеснок. Шаг 5: Добавить томатную пасту, немного воды и тушить до готовности.',
 5.2, 12, 7.1, 141.6, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/85562.jpg', 'e48e6f63bd8074965219003148e3229b99bf48fc57accf0835cd3993a75ca395'),

-- https://calorizator.ru/recipes/59054
( 'Бургер по-домашнему',
 'Шаг 1: Смолоть овсяные отруби в блендере. Шаг 2: Смешать сухие ингредиенты для булочки. Шаг 3: Взбить блендером яйцо и воду. Шаг 4: Добавить сухие ингредиенты и еще раз хорошо взбить. Шаг 5: Сформировать два шара и приплюснуть их. Шаг 6: Выложить на противень, застеленный пекарской бумагой/силиконовым ковриком. Выпекать в духовке 30 минут. Шаг 7: Фарш смешать с яйцом. Шаг 8: Сформировать котлетки и обжарить на антипригарной сковородке в течение 6 минут. Шаг 9: Смешать йогурт и томатную пасту. Шаг 10: Порезать помидор на кружочки. Шаг 11: Сформировать бургер: разрезать булочку на две части, нижнюю намазать соусом, положить лист салата, сверху помидор, котлету, помидор, лист салата, верхнюю часть помазать соусом и закрыть бургер.',
 10.8, 8.2, 10.3, 170.9, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/59054.jpg', '42462b43e3a6d79b838c356b933a09a5c99c8fb3c3b27f5ff836df3d20d66b45'),

-- https://calorizator.ru/recipes/102228
( 'Греча с курицей в духовке',
 'Шаг 1: Промыть гречку. Шаг 2:Смазать форму каплей масла. Шаг 3:Выложить в форму промытую гречку. Шаг 4:Выложить курицу поверх гречки. Шаг 5:Разложить овощи поверх гречки с курицей, залить водой. Шаг 6:Накрыть фольгой и отправить в духовку на 30 минут при 180 градусах. Шаг 7:Снять фольгу и поддержать в духовке еще 15 минут, можно включить режим гриля.',
 11.3, 8.6, 6.3, 134.8, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/102228.jpg', 'be750f5708b38aef9c30495044dc0bf1201a15e9b72239e3edda063fae357e56'),

-- https://calorizator.ru/recipes/82050
( 'Гуляш рыбный',
 'Шаг 1: Лук мелко нарезать, обжарить на антипригарной сковороде до золотистого цвета, помешивая ( минуты), на среднем огне. Шаг 2: Перец нарезать небольшими кубиками, добавить к луку, обжарить 3-4 минуты. Шаг 3: Лук и перец залить протертыми томатами, уменьшить огонь до минимума, накрыть крышкой и тушить 10 минут. Шаг 4: Треску нарезать произвольно, выложить в томатно-перечный соус и тушить под крышкой 10 минут. Шаг 5: Креветки разморозить, очистить, добавить к рыбе и тушить еще 5 минут. Шаг 6: Из оливок слить жидкость, добавить в гуляш, приправить солью и перцем, тушить все вместе еще 5 минут.',
 10.7, 2.4, 1.4, 65.3, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/82050.jpg', '123793faa58a2749a2c417b57b3934737d776c5590e99518ddc9b28e32a5393c'),

-- https://calorizator.ru/recipes/58547
( 'Драники с мясом',
 'Шаг 1: Подготовить фарш. Шаг 2: Картофель очистить и пропустить через мясорубку. Шаг 3: Лук пропустить через мясорубку и добавить к картофелю. Шаг 4: К картофелю добавить фарш и яйцо. Все посолить и поперчить. Хорошо перемешать. Шаг 5: К ингредиентам добавить муку и еще раз перемешать. Шаг 6: Столовой ложкой выложить подготовленную смесь на раскалённую сковороду. Шаг 7: Драники с мясом обжарить с обеих сторон до золотистой корочки.',
 5.7, 13.7, 15.1, 211.7, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/58547.jpg', 'da1f3169ce79de0a8ea62f9153ec5955ae594e234e09fb8b7cc2401f3893baf2'),

-- https://calorizator.ru/recipes/67103
( 'Жареные патиссоны',
 'Шаг 1: Мелко нарезать лук. Шаг 2: Обжарить лук в растительном масле. Шаг 3: Небольшими кубиками нарезать патиссоны. Шаг 4: Добавить патиссоны к луку и жарить помешивая. Шаг 5: Мелкими кубиками нарезать морковь. Шаг 6: Добавить морковь к патиссонам размешать, посолить и жарить минут 15 до тех пор, пока морковь и патиссоны не станут мягкими. Шаг 7: Мелко нарезать укроп, чеснок выдавить через пресс. Шаг 8: В конце приготовления добавить укроп с чесноком, перемешать и прожарить 1-2 минуты.',
 0.9, 5.8, 6.4, 82.6, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/67103.jpg', '23453626225a81c76c57dc4445d23ed1ffc2ad9ba5b99389dff4c642454610d6'),

-- https://calorizator.ru/recipes/136443
( 'Кальмар с грибами в сметане',
 'Шаг 1: Лук нарезать полукольцами. Шаг 2: Нарезать шампиньоны. Шаг 3: Ошпарить кальмары горячей водой, чтобы свернулась верхняя пленка. Шаг 4: Промыть кальмары под струей холодной воды, удалить верхнюю пленку, внутренности и позвоночник. Шаг 5: Нарезать кальмаров кольцами. Шаг 6: В сковороду влить масло, выложить лук и грибы, пассеровать 5-7 минут. Шаг 7: Добавить нарезанных кальмаров, тушить 3 минуты под крышкой.. Шаг 8: Добавить сметану, подсолить, перемешать и тушить еще минуту.',
 10.2, 1.8, 2.8, 70.1, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/136443.jpg', '52c3f04fc04009a51d03cee51a58bf4912fdc68092897b7019ca9acbd925d9e7'),

-- https://calorizator.ru/recipes/103641
( 'Капуста тушеная с соевым гуляшом',
 'Шаг 1: Соевый гуляш замочить в подсоленной теплой воде на 20 минут. Шаг 2: Капусту нашинковать, морковь натереть на терке. Шаг 3: На сковороде разогреть масло и специи, примерно полминуты.  Шаг 4: Пассеровать морковь на масле со специями. Шаг 5: Добавить соевый гуляш в сковороду, тушить несколько минут, периодически помешивая. Шаг 6: Добавить капусту и долить 50 мл. воды, подсолить, тушить до готовности. Шаг 7: Добавить к яйцам укроп и приправы, перемешать.',
 3.7, 4.8, 2.1, 54.2, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/103641.jpg', 'e8446a0a0d4f4b9a8ff67ec8a227c8c01b896ffeafd152b6d860060359d89b4a'),

-- https://calorizator.ru/recipes/37887
( 'Картофельный салат с маслинами',
 'Шаг 1: Картофель помыть, почистить и отварить. Нарезать кубиками, но не мельчить, иначе получится однородная масса. Шаг 2: Красный лук почистить и нашинковать полукольцами или четвертинками. Вместо красного лука можно использовать белый салатный или ялтинский сорт лука. Шаг 3: Маслины порезать вдоль пополам. Шаг 4: Огурец порезать мелкими кубиками. Вместо соленого огурца можно использовать маринованные огурцы или корнишоны. Смешать все ингредиенты. Шаг 5: Смешать в отдельной чашке оливковое масло и горчицу, выдавить сок четверти лимона, всё хорошенько взбить вилкой. Горчицу лучше использовать не острую. Вместо лимона можно использовать лайм. Шаг 6: Зелень петрушки измельчить. Добавить к овощам, все посолить, поперчить, полить заправкой и хорошо перемешать. Шаг 7: Добавить к яйцам укроп и приправы, перемешать.',
 1.8, 11.4, 10.2, 149, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/37887.jpg', '38f3bfdcd6d61d6735c38a140e48595aa17fa50b672d09c1490cf8a4fb9d5a0e'),

-- https://calorizator.ru/recipes/16485
( 'Копченые грудки по-домашнему',
 'Шаг 1: Куриные грудки разрезать на две части. Шаг 2: Имбирь натереть, смешать в кастрюле с вином, соевым соусом и оливковым маслом. Шаг 3: Замочить в соевом маринаде куриные грудки и поставить в холодильник на 2-3 часа. Шаг 4: После маринования кусочки обсушить и натереть черным молотым перцем и солью, оставить на 1 час при комнатной температуре. Шаг 5: Обжарить грудки на сильном огне по 3 минуты с каждой стороны. Шаг 6: На дно большой сковороды выложить фольгу. Смешать рис, чай, корицу и выложить на фольгу. Поставить сковороду в духовку, сверху поместить решетку. Шаг 7: Выложить грудки на решетку, поставить сковороду на сильный огонь и прогреть 7 минут, не закрывая крышки. Шаг 8: Затем накрыть крышкой филе готовить 10–20 минут на среднем огне.',
 20.1, 7.9, 6.1, 167.7, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/16485.jpg', '23d679d0c8d5ad86f54e0258fd5d1323495a77044e0440423979f035ef077532'),

-- https://calorizator.ru/recipes/87004
( 'Tuna salad по-вегетариански',
 'Шаг 1: Нут отварить до готовности. Шаг 2: Огурцы нарезать квадратиками. Шаг 3: Авокадо размять вилкой. Шаг 4: Нори нарезать квадратиками. Шаг 5: Измельчить нут до состояния пюре (но с небольшими кусочками). Шаг 6: Соединить все ингредиенты. Шаг 7: Можно подавать как салат, можно мазать на бутерброд, можно делать тарталетки. Шаг 8: Грибы выложить на дно формы смазанной маслом.',
 4.1, 12, 4.3, 101.1, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/87004.jpg', '22f20917e6a52174bae0f05b04de7c0d9f6a8e3a8c558d9d2f87c0e8ad4ad28f'),

-- https://calorizator.ru/recipes/118622
( 'Беляши из кабачков с фаршем',
 'Шаг 1: Кабачок натереть на терке, подсолить, оставить на несколько минут, затем слить жидкость. Шаг 2: Лук мелко нарезать, петрушку мелко порубить ножом. Шаг 3: Добавить муку к кабачкам и замесить тесто. Шаг 4: Фарш обжарить с луком. Шаг 5: На сухую сковороду выложить лепешки из кабачкового теста, в центр каждой выложить фарш, затем снова тесто. Шаг 6: Обжарить с двух сторон на сильном огне.',
 8.1, 8.3, 3.6, 101.6, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/118622.jpg', '161b23b7b9ca6f529679267c3e0f3622bf25d869744923d1c41b425ba42f2d9f'),

-- https://calorizator.ru/recipes/106821
( 'Бутерброды с нутом и черри',
 'Шаг 1: Нут сильно разварить в подсоленной воде, лишнюю воду слить. Шаг 2: Черри нарезать на половинки или четвертинки. Шаг 3: Вареный нут переложить в тарелку и, пока он горячий, добавить в него специи. Шаг 4: Размять вилкой нут в пюре. Шаг 5: На каждый ломтик хлеба выложить пасту из нута. Шаг 6: Выложить кусочки помидоров черри. Шаг 7: Добавить к яйцам укроп и приправы, перемешать.',
 5.1, 23.6, 1.2, 126.9, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/106821.jpg', '41bdbcfa70624cab2ec3069e4e1bceeea38b4a332784d731d23f576da01e5955'),

-- https://calorizator.ru/recipes/81613
( 'Галета ржаная с томатами',
 'Шаг 1: Просеянную муку (250 гр.) смешать с томатным соком, яйцом, солью и чёрным перцем. Шаг 2: Скатать тесто в шар и оставить на 10 минут. Шаг 3: Круглую форму для выпекания присыпать мукой (гр.) и раскатать тесто, сделав бортики. Шаг 4: Сыр нарезать кружками, а кружки – пополам. Шаг 5: Нарезать помидоры кружками и пополам. Шаг 6: Сыр и помидоры, чередуя, выложить на середину теста и присыпать травами. Шаг 7: Бортики опустить на начинку. Шаг 8: Выпекать галету в разогретой до 180 °С духовке 40 минут.',
 7, 21.4, 6.1, 157.4, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/81613.jpg', '4ed2c1eca48ddc8de407386e529e0ac71d77c818a3ea5b22a2f53ca4202dbbc5'),

-- https://calorizator.ru/recipes/76752
( 'Грибы фаршированные творогом',
 'Шаг 1: Творог смешать с яйцом и паприкой. Шаг 2: Мелко нарезать петрушку. Шаг 3: Добавить зелень к творожно-яичной массе. Шаг 4: У грибов удалить ножки. Шаг 5: Поместить на противень и запекать при 200 °С 20 минут.',
 9.8, 1.1, 3.9, 80.9, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/76752.jpg', 'b07a3675f3342edcbe987bde4feb36ec8d29eab1255d4d4ceddcebe53b0ff37e'),

-- https://calorizator.ru/recipes/171223
( 'Густой суп с чечевицей и булгуром',
 'Шаг 1: Нарезать картофель кубиками. Шаг 2: Мелко нашинковать лук. Шаг 3: Натереть морковь на крупной терке. Шаг 4: Бульон развести водой, довести до кипения, добавить картофель, чечевицу и булгур. Варить 20 минут. Шаг 5: В сковороду влить масло, пассеровать лук 3 минуты. Шаг 6: Добавить морковь к луку, пассеровать еще пять минут. Шаг 7: Добавить к овощам томатную пасту, чеснок и сахар, перемешать. Шаг 8: Влить в сковороду немного бульона из кастрюли и тушить 10 минут. Шаг 9: Добавить зажарку в суп. Шаг 10: За пару минут до готовности добавить в суп отварное куриное мясо.',
 1.9, 4, 0.9, 32.7, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/171223.jpg', '75ca49c988bef44a446712d3b7015adb12f4d4429134b0f32bffccf778a42013'),

-- https://calorizator.ru/recipes/79165
( 'Закуска из творога и овощей',
 'Шаг 1: Кабачок нарезать на кружочки толщиной 3-4 см. Шаг 2: Ложкой извлечь мякоть, но не до конца. Шаг 3: Мелко нарезать укроп и чеснок. Шаг 4: Морковь и серединки кабачка натереть на мелкой терке. Шаг 5: Творог смешать с измельченными овощами, зеленью и чесноком. Шаг 6: Начинить смесью кружочки кабачка. Шаг 7: Запекать в духовке при 180 °С 20-25 минут.',
 2.4, 5, 0.3, 32.5, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/79165.jpg', 'e750d82908c7c4bea55ae414791a9c20bf19f16315b1bcf1a5d681feef917c51'),

-- https://calorizator.ru/recipes/52770
( 'Запеканка с рисом, сыром и колбасой',
 'Шаг 1: Яйца отварить в воде вкрутую 10 мин. Остудить в холодной воде и очистить. Шаг 2: Рис хорошо промыть. Отварить в подсоленной воде до готовности. Шаг 3: Вареную колбасу нарезать кубиками. Шаг 4: Яйца тоже нарезать кубиками. Шаг 5: Помидоры нарезать кружочками, потом разрезать кружки пополам. Шаг 6: Натереть сыр на крупной терке. Шаг 7: Рис, яйца, колбасу и 2/3 сыра сложить в миску, поперчить и хорошо перемешать. Шаг 8: Форму смазать растительным маслом. Выложить в нее ровным слоем подготовленные продукты. Сверху смазать сметаной. Шаг 9: На сметану выложить помидоры. Разогреть заранее духовку до 180 градусов. Поставить форму в духовку и запекать 15 минут. За 5 минут до конца запекания посыпать блюдо оставшимся сыром. Подавать запеканку горячей. Перед подачей украсить зеленью.',
 8.3, 19.6, 11.3, 213.7, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/52770.jpg', '7b7f5b2f6ca8c5181741ed90a56e5dc5f3d138a79f27fdc5023c1490c7cb3a77'),

-- https://calorizator.ru/recipes/33554
( 'Запеченые лодочки',
 'Шаг 1: Из каждой булочки аккуратно извлечь мякоть. Шаг 2: Бекон обжарить на сковороде, мелко порезать. Шаг 3: Сыр натереть на терке, лук мелко порубить. Шаг 4: Сыр натереть на терке, лук мелко порубить. Шаг 5: К яйцам добавить бекон, лук и сыр, перемешать. По вкусу добавить соль, перец. Шаг 6: Приготовленную начинку уложить в лодочки из булочек. Шаг 7: Поместить булки на противень, застеленный пергаментом и отправить в разогретую до 175 градусов духовку на 20—25 минут.',
 10.4, 33.4, 10.5, 271.7, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/33554.jpg', '862b182e42ed487c99e9601391499e3473ef6c47f7b9c6ebc42e6f12b9252a2a'),

-- https://calorizator.ru/recipes/33557
( 'Куриные рулетики с ветчиной и сыром',
 'Шаг 1: Куриные грудки промыть и хорошенько отбить. Шаг 2: Ветчину нарезать длинными пластами и поместить на отбитые грудки. Шаг 3: Поверх ветчины поместить по две пластины плавленого сыра. Шаг 4: Голубой сыр нарезать соломкой и поместить несколько соломок с краю отбитого филе. Шаг 5: Перец промыть и нарезать соломкой, поместить по несколько соломой поверх голубого сыра. Шаг 6: Свернуть каждое филе в рулетик и закрепить зубочистками. Шаг 7: Уложить рулетики в смазанную оливковым маслом форму для выпечки швом вниз. Шаг 8: В чаше смешать оливковое масло с солью и смазать рулетики. Шаг 9: Отправить в предварительно разогретую до 200 градусов духовку на 30 минут. Перед подачей на стол убрать зубочистки и нарезать рулетики.',
 20.3, 2, 6.2, 149.1, 'https://calorizator.ru/sites/default/files/imagecache/recipes_full/recipes/33557.jpg', '9e4679540b36fb4e6dfe8659955f53d8b69f02360e23ea40dd855c906697cb65');
