# Details

Date : 2024-06-10 18:55:50

Directory c:\\Users\\vibeb\\Desktop\\Диплом\\разработка\\osau_ams_api

Total : 90 files,  8985 codes, 1830 comments, 1870 blanks, all 12685 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [cmd/app/main.go](/cmd/app/main.go) | Go | 8 | 18 | 10 | 36 |
| [configs/config.yml](/configs/config.yml) | YAML | 16 | 1 | 2 | 19 |
| [domain/attendance.go](/domain/attendance.go) | Go | 44 | 0 | 9 | 53 |
| [domain/classroom.go](/domain/classroom.go) | Go | 5 | 0 | 2 | 7 |
| [domain/departament.go](/domain/departament.go) | Go | 17 | 0 | 4 | 21 |
| [domain/discipline.go](/domain/discipline.go) | Go | 13 | 0 | 4 | 17 |
| [domain/disciplineType.go](/domain/disciplineType.go) | Go | 5 | 0 | 2 | 7 |
| [domain/educationLevel.go](/domain/educationLevel.go) | Go | 5 | 0 | 2 | 7 |
| [domain/educationType.go](/domain/educationType.go) | Go | 5 | 0 | 2 | 7 |
| [domain/faculty.go](/domain/faculty.go) | Go | 17 | 0 | 4 | 21 |
| [domain/group.go](/domain/group.go) | Go | 12 | 0 | 4 | 16 |
| [domain/headman.go](/domain/headman.go) | Go | 14 | 0 | 4 | 18 |
| [domain/profile.go](/domain/profile.go) | Go | 14 | 0 | 4 | 18 |
| [domain/report.go](/domain/report.go) | Go | 39 | 0 | 5 | 44 |
| [domain/schedule.go](/domain/schedule.go) | Go | 26 | 0 | 5 | 31 |
| [domain/specialty.go](/domain/specialty.go) | Go | 15 | 0 | 4 | 19 |
| [domain/student.go](/domain/student.go) | Go | 13 | 0 | 3 | 16 |
| [domain/teacher.go](/domain/teacher.go) | Go | 21 | 0 | 5 | 26 |
| [domain/university.go](/domain/university.go) | Go | 9 | 0 | 2 | 11 |
| [domain/user.go](/domain/user.go) | Go | 20 | 0 | 5 | 25 |
| [internal/app/app.go](/internal/app/app.go) | Go | 77 | 7 | 14 | 98 |
| [internal/config/config.go](/internal/config/config.go) | Go | 62 | 3 | 15 | 80 |
| [internal/handler/attendance.go](/internal/handler/attendance.go) | Go | 306 | 118 | 65 | 489 |
| [internal/handler/classroom.go](/internal/handler/classroom.go) | Go | 114 | 64 | 29 | 207 |
| [internal/handler/departament.go](/internal/handler/departament.go) | Go | 166 | 86 | 35 | 287 |
| [internal/handler/discipline.go](/internal/handler/discipline.go) | Go | 142 | 86 | 35 | 263 |
| [internal/handler/disciplineType.go](/internal/handler/disciplineType.go) | Go | 114 | 64 | 29 | 207 |
| [internal/handler/educationLevel.go](/internal/handler/educationLevel.go) | Go | 114 | 64 | 29 | 207 |
| [internal/handler/educationType.go](/internal/handler/educationType.go) | Go | 123 | 75 | 32 | 230 |
| [internal/handler/errors.go](/internal/handler/errors.go) | Go | 2 | 0 | 2 | 4 |
| [internal/handler/faculty.go](/internal/handler/faculty.go) | Go | 170 | 98 | 35 | 303 |
| [internal/handler/group.go](/internal/handler/group.go) | Go | 130 | 86 | 35 | 251 |
| [internal/handler/handler.go](/internal/handler/handler.go) | Go | 220 | 5 | 36 | 261 |
| [internal/handler/headman.go](/internal/handler/headman.go) | Go | 133 | 78 | 32 | 243 |
| [internal/handler/middleware.go](/internal/handler/middleware.go) | Go | 64 | 0 | 11 | 75 |
| [internal/handler/profile.go](/internal/handler/profile.go) | Go | 157 | 97 | 38 | 292 |
| [internal/handler/regexes.go](/internal/handler/regexes.go) | Go | 20 | 0 | 4 | 24 |
| [internal/handler/reponse.go](/internal/handler/reponse.go) | Go | 21 | 2 | 8 | 31 |
| [internal/handler/report.go](/internal/handler/report.go) | Go | 44 | 0 | 11 | 55 |
| [internal/handler/schedule.go](/internal/handler/schedule.go) | Go | 419 | 212 | 87 | 718 |
| [internal/handler/specialty.go](/internal/handler/specialty.go) | Go | 142 | 86 | 35 | 263 |
| [internal/handler/student.go](/internal/handler/student.go) | Go | 141 | 75 | 32 | 248 |
| [internal/handler/teacher.go](/internal/handler/teacher.go) | Go | 160 | 86 | 35 | 281 |
| [internal/handler/university.go](/internal/handler/university.go) | Go | 151 | 82 | 32 | 265 |
| [internal/handler/users.go](/internal/handler/users.go) | Go | 352 | 154 | 55 | 561 |
| [internal/handler/validators.go](/internal/handler/validators.go) | Go | 55 | 0 | 19 | 74 |
| [internal/repository/attendance.go](/internal/repository/attendance.go) | Go | 223 | 16 | 39 | 278 |
| [internal/repository/classroom.go](/internal/repository/classroom.go) | Go | 87 | 20 | 29 | 136 |
| [internal/repository/departament.go](/internal/repository/departament.go) | Go | 195 | 0 | 37 | 232 |
| [internal/repository/discipline.go](/internal/repository/discipline.go) | Go | 179 | 4 | 37 | 220 |
| [internal/repository/disciplinetype.go](/internal/repository/disciplinetype.go) | Go | 103 | 4 | 29 | 136 |
| [internal/repository/educationlevel.go](/internal/repository/educationlevel.go) | Go | 103 | 4 | 29 | 136 |
| [internal/repository/educationtype.go](/internal/repository/educationtype.go) | Go | 107 | 0 | 29 | 136 |
| [internal/repository/faculty.go](/internal/repository/faculty.go) | Go | 195 | 0 | 37 | 232 |
| [internal/repository/group.go](/internal/repository/group.go) | Go | 175 | 8 | 37 | 220 |
| [internal/repository/headman.go](/internal/repository/headman.go) | Go | 147 | 4 | 29 | 180 |
| [internal/repository/profile.go](/internal/repository/profile.go) | Go | 234 | 12 | 45 | 291 |
| [internal/repository/report.go](/internal/repository/report.go) | Go | 150 | 0 | 8 | 158 |
| [internal/repository/repository.go](/internal/repository/repository.go) | Go | 228 | 0 | 23 | 251 |
| [internal/repository/schedule.go](/internal/repository/schedule.go) | Go | 576 | 1 | 99 | 676 |
| [internal/repository/specialty.go](/internal/repository/specialty.go) | Go | 179 | 24 | 37 | 240 |
| [internal/repository/student.go](/internal/repository/student.go) | Go | 140 | 24 | 37 | 201 |
| [internal/repository/teacher.go](/internal/repository/teacher.go) | Go | 175 | 16 | 37 | 228 |
| [internal/repository/university.go](/internal/repository/university.go) | Go | 84 | 16 | 28 | 128 |
| [internal/repository/users.go](/internal/repository/users.go) | Go | 570 | 10 | 55 | 635 |
| [internal/server/server.go](/internal/server/server.go) | Go | 26 | 0 | 7 | 33 |
| [internal/service/attendance.go](/internal/service/attendance.go) | Go | 62 | 0 | 13 | 75 |
| [internal/service/classroom.go](/internal/service/classroom.go) | Go | 37 | 0 | 11 | 48 |
| [internal/service/departament.go](/internal/service/departament.go) | Go | 58 | 0 | 13 | 71 |
| [internal/service/discipline.go](/internal/service/discipline.go) | Go | 46 | 0 | 13 | 59 |
| [internal/service/disciplinetype.go](/internal/service/disciplinetype.go) | Go | 37 | 0 | 11 | 48 |
| [internal/service/educationlevel.go](/internal/service/educationlevel.go) | Go | 37 | 0 | 11 | 48 |
| [internal/service/educationtype.go](/internal/service/educationtype.go) | Go | 40 | 0 | 12 | 52 |
| [internal/service/errors.go](/internal/service/errors.go) | Go | 10 | 0 | 3 | 13 |
| [internal/service/faculty.go](/internal/service/faculty.go) | Go | 58 | 0 | 13 | 71 |
| [internal/service/group.go](/internal/service/group.go) | Go | 43 | 0 | 13 | 56 |
| [internal/service/headman.go](/internal/service/headman.go) | Go | 43 | 0 | 12 | 55 |
| [internal/service/profile.go](/internal/service/profile.go) | Go | 52 | 0 | 14 | 66 |
| [internal/service/report.go](/internal/service/report.go) | Go | 16 | 0 | 6 | 22 |
| [internal/service/roles.go](/internal/service/roles.go) | Go | 5 | 0 | 2 | 7 |
| [internal/service/schedule.go](/internal/service/schedule.go) | Go | 103 | 0 | 23 | 126 |
| [internal/service/service.go](/internal/service/service.go) | Go | 76 | 1 | 9 | 86 |
| [internal/service/specialty.go](/internal/service/specialty.go) | Go | 49 | 0 | 13 | 62 |
| [internal/service/student.go](/internal/service/student.go) | Go | 52 | 0 | 13 | 65 |
| [internal/service/teacher.go](/internal/service/teacher.go) | Go | 55 | 0 | 13 | 68 |
| [internal/service/university.go](/internal/service/university.go) | Go | 52 | 0 | 12 | 64 |
| [internal/service/users.go](/internal/service/users.go) | Go | 162 | 0 | 27 | 189 |
| [pkg/auth/jwtManager.go](/pkg/auth/jwtManager.go) | Go | 47 | 0 | 14 | 61 |
| [pkg/database/postgres/postgres.go](/pkg/database/postgres/postgres.go) | Go | 24 | 16 | 10 | 50 |
| [pkg/myhash/password.go](/pkg/myhash/password.go) | Go | 28 | 3 | 9 | 40 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)