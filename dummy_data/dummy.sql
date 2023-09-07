INSERT INTO `members` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `project_id`) VALUES
(1, '2023-09-05 22:25:02.992', '2023-09-05 22:25:02.992', NULL, 'Rahadian', 'Project Manager', 1),
(2, '2023-09-05 22:25:02.992', '2023-09-05 22:25:02.992', NULL, 'Ardya', 'Backend Engineer', 1),
(3, '2023-09-05 22:25:02.992', '2023-09-05 22:25:02.992', NULL, 'Koto', 'Frontend Engineer', 1),
(4, '2023-09-05 22:25:02.992', '2023-09-05 22:25:02.992', NULL, 'Panjang', 'Frontend Engineer', 1),
(5, '2023-09-05 22:25:53.156', '2023-09-05 22:25:53.156', NULL, 'TEJO', 'Project Manager', 2),
(6, '2023-09-05 22:25:53.156', '2023-09-05 22:25:53.156', NULL, 'GUNTUR', 'Backend Engineer', 2),
(7, NULL, NULL, NULL, 'Surya', 'Project Manager', 3);


INSERT INTO `projects` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `client_name`, `budget`, `progress`) VALUES
(1, '2023-09-05 22:25:02.991', '2023-09-05 22:25:02.991', NULL, 'Project 1', 'PERTAMINA', 50000000, NULL),
(2, '2023-09-05 22:25:53.155', '2023-09-07 11:46:05.436', NULL, 'Project 2', 'SAMPOERNA', 150000000, 0.1),
(3, '2023-09-07 09:45:29.472', '2023-09-07 09:45:29.472', '2023-09-07 10:27:08.173', 'Project 3', 'PETROSEA', 10000000, 0);
