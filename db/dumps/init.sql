CREATE TABLE IF NOT EXISTS movies
(
    id           SERIAL PRIMARY KEY,
    title         varchar(100),
    description  text,
    image text,
    updated_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT not_empty_title CHECK (title <> '')
);

CREATE UNIQUE INDEX IF NOT EXISTS movies_title_idx ON movies (title);
CREATE INDEX IF NOT EXISTS movie_description_idx ON movies (description);


CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL PRIMARY KEY,
    username     varchar(255),
    password     varchar(255),
    role         varchar(100),
    updated_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT not_empty_username CHECK (username <> ''),
    CONSTRAINT not_empty_password CHECK (password <> '')
);

CREATE UNIQUE INDEX IF NOT EXISTS users_username_idx ON users (username);
CREATE INDEX IF NOT EXISTS users_password_idx ON users (password);

CREATE TABLE IF NOT EXISTS comments
(
    id           SERIAL PRIMARY KEY,
    content      text NOT NULL,
    user_id bigint REFERENCES users (id) ON DELETE CASCADE,
    movie_id bigint REFERENCES movies (id) ON DELETE CASCADE,
    created_at   timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT not_empty_content CHECK (content <> '')
    );

CREATE TABLE IF NOT EXISTS states
(
    state_id uuid,
    user_id bigint REFERENCES users (id) ON DELETE CASCADE,
    password varchar(255),
    code bigint,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS states_id_idx ON states (state_id,code,user_id);

-- Добавление тестовых данных
INSERT INTO users
    (username, password, role)
VALUES ('molka4477@gmail.com', '123', 'user'),
       ('smilee88@yandex.ru', 'qwerty', 'user');

INSERT INTO movies
(title, description, image)
VALUES ('Джанго', 'Как всегда, Дикий Запад наполнен бандитами и головорезами. Терроризм процветает на территории небольшого городка, пока в нем не появляется защитник слабых и угнетенных Джанго. Он появляется ниоткуда, притащив за собой на веревке гроб с «секретом». Трудна миссия по защите угнетенных, но, как всегда, выполнима, тем более, если за дело берется Джанго.', 'https://kinopoisk-ru.clstorage.net/2cN9w8350/dbcad21QyAj/ZLTywhm6vTuulVYDpQLu0uTGFm-McbyCgNpEscudfabVO--cl4oZvsCgjuruznoDWilZhUlDnWvINZC35sXzM_Fkp1XdPeqveQdF1vCXe4DiRAI8i-OpEYaMgP2HkTYsHjLwVJ0AE_pWsQj8utkvk0o_iwcbUszyUzbPYqNSbm6hMDZKmKQ3eCwDEwYtgWESfMRWo1E9NQRfHcP4jJo37tA9vhvciBY82vCfNjxX4NwK-4tACZ-FNlQnRiOrmqqtKLE4Bhxpd_7nxRQKYsntWPYA3i-bsPTKlpgMsYHbtelZvX4V1QDUdJ_0Gjx2mOTbi_BDR4KcCX-S7kQhrdMi7a4y8wVe4fKxY4PWSqKAIt79RN0n2zVwTVaSQHPCGj6hEbu5XNwAUPQWOdg8Md4nEoU3RIPLl4P_XqJNaSuR4-ThPbbKFei_-CRK00-kSiJf84zRrNu898pbWka-i1x4bh__eZacjhm32n2VOHkXr98M8cbMBF8BNlxiBSEglmZq4DXxx5gvvXZhRRpNb4um0buA3OmSOLcE1RHKe8pd-qkTfHhVVEueNtnwGrjzWOVUgvTNw0vVADvSpAZrIxJmK620doeX4XQ3p0dWCKSOadc7CV0mXPu-xZnegrPFXPLqm3F3lN0BWTnWOl5zvFDqWsg-gsbL3Yi1V61JYCpVqquie_kGX-T3fC5Kkg-nSG0Vds4TJ5fxP0OX0Uw-QNK05Bex9pdcQJ493jdSNDIX4F4IM0TPg1sPs9frBSRt3W1n5nm3xtAl8rDpSliGYIKsXTxBXOpSfHSFUl2DOYUc_q4Qez6QWMOadBj0FX0-mi1UC3_FiENcC34YIM6rL9BqrC799gWXoXh6aQJVD2jH51g2BFznGfS8hlsbC3vL0zclnbm-VZGP2fuQ8BjyctJoUUr8xQYF0MHxV-hB4GVeqKqq9XnOXSIysSoGkUGrwikdsgdXaNX6PgIeXY13DZx1ZFUyvl-XQtR_U_UWdXXa7huEcMwBTp1PNpAnhCTsnOkgZLe7jZYsO3_vC1vLJcolkXpPUiaX_vjC1x2PcsZRMqfRPH-eXkPZ9h_1kjt9kC5aDTdLB8naQH-dqsQmbtHoLqE6uUbRbzg-K4wUiizAJV-yQVlumbH-gRMRwvBB0LglWPt92FQDnDQdcNY2PB7tms07goODkMs1F6pOLmDcpSWgtfcEl68ztKvGk09gCy6VvIbfrpqydwcQ0E13jFw-qZI_uRtZTpN2VTbecrlTZRuDt8rPgB7LN59qhaFiEKAvpXw_DxkscPdqwJHNZcYvHLXPVKtW8PkFEJ6ONsNbPOhVtHSb2o-U9Rb_kL613mUdQnFDA4OfhfHVYglgYBPsJ-J08IrcqD8y7gPcg2cLrJZ_CdSunHy1BlfTAjRLlTbv1Ti3mpnEWDSb_d729NosHgv_wwHJ34r8GqtCoGsZaqCg_DpDn6N-_ugMnkAkDqofPM6W6NZ190pY0gWyjtm_ZZp0uNsXiRg6kfCWNjRY758AuE4MwhxA917tRy9g3SLgZDNzjx-oMzLoxZuOpcXqmTKJGipUsXaL1RUA-MnQ-2AT8LKTGkTVeZH50PVwHiLRB3iOhYlVTbRRb0-uKtHsIqV6v8Yd5bO5ZsQTzqDIohi3Bpws3vS1i1tYBL8KEfhtGvH0UxGG0PWef1B7PBEg1U35joGNl8E0H-THp2XerKIl_LuPXiN2sShKE4ovh65V_EaTrx21NQtSEwe3jNM7IZfzflpXDx33WHSf-7uT5l2KecFHwVYMvxXsjKiiUeJlLnm2y9jjdHZpDduNoI2ml_OB0-Je-XwHWl3DNAsbtGfbuPyb1A-cNxr4H3y71qbWgLPJTo7QyLTUbYdnbxrsIGK99k5bZbKyIIzUyaPGL9q0wNyg0ny4gRWexbdDlXIgWziwHdxJ0rpafxEz-ldi3kT3Ao2K3gr2UOWKZO7ZYucm8rCN02z18mxCHULvAmyRewjeql52_M3RnI57SBP3bFuwMRzdCtw22PJWev1Y6VQFvEvHSB8B8FhqyWzn2KluqfP2xxEqf3-uhNgOq8mu2vrJWSEWdDWGkBBD_4PQuCRfvbwYlQPbMt1_0vh7keQcgHBJiMLShbMUZghqJdyjIip0uERfIbc4os6ZAuVPJhD7xNdv3fX_xhgUjr5G1fet33RxmlJClv4Q9163eNwh1QC5C4fBUsB-keUI520a6aZg-P-NkKw3eagHnECqCCuS-wnT6V97tMNYGsz0yx22qtB5uZ1eSNw6Ez7X8D6XYNQNu8xOi12B8VHvgqmtmu1l4n25QpUp83anBhyJ7IEjHnzAVuoXdbjF29QNswAYfS8Uu3wSFY4VtJ68WzeyHiXRQL_NSEwaDDFR4Ejv6J7o7yyycAxZqbR2qYRVSKxHal87QNUgGvj_wp-VgvAFU7AvGLOwFp8OWXUY8FZ0dpYhmog-TIiMX0s1XCrIqCRS6aQlf3JMFap8dK5M1Q7uSK0VfAAXqNUwMcGbU4oxyVJ_aJR0vBIfihG1E_mb97PR5V4Me8THTNND9dKugCfjneMs5bz8zJ9s_vApQ5VOY0LnVLrIUmDdNTxM1lwAfwlStSUbsPEVEIcafRIx2bW2l-7TwntOQUhXw7-RJQgr4hoiauWxMYwQa330L4tZDqcIbtR6DNOrXjO-A1WeBftIF_2lUb361h3Hk7GffxKwcRDg1YU2gYHBm4P_VKvOoarVYuXhdHGF2Gv6euZD2w-tBeZSvIycadmx-c4Rm482TNP95ZG8cpzfiJq0n7xZej1cJVYBNI-FRJUL-BwtgGBlma1v6Pj5hxdi9X9oThMN6kHuF3zD02rW-75CV1bCMIkdv-8Rc3Qf1gCec5J52LS-nOAZyPdMj40XyTlRagZkJFWlamyyckXR67D3bMNQBq5B6xDxTR1pnLO4jNPTy3BKWvLhUv2xlp2HGbKQuNY_ORCoFAr5Do6Aloe8EY'),
       ('Святые из Бундока', 'Чего только не бывает на свете! Два обычных ирландских парня, братья Коннор и Мерфи, преспокойно жили и работали в своем родном Бостоне, пока в один прекрасный день на них не снизошло озарение: сам Бог послал их на землю с особой миссией, наделив святой силой, чтобы очистить мир от зла. И братья взялись за дело со всей серьезностью, присущей ирландцам...', 'https://kinopoisk-ru.clstorage.net/2cN9w8350/dbcad21QyAj/ZLTywhm6vTuulVYDpQLu0uTGFm-McbyCgNpEscudfabVO--cl4oZvsChT-psz7hDGiibk57AXGnJ9ZHi54WzM-XlZxXdKCv6eRLElnLDboA3E8N_iuOoU4eekKZcVrGs2bGzFRhAEbXY8R67qxhuFMo5zAeMFJo5GakLMm0QJaRq-bdHWaU7t6NEVgYlxiSWecMcq145s0rbW0p3DVO6p9hzfRpWD1A93TeS-jzS6JoFdozDjFOI9FBtDCPkUS1srjo-xlfpv3rjwp3AIspr3vWJlm_WtXnHGpWD8sIat-ZadPSalgFZtR9_1_652KJVAXENyERUA_fcJMrq5lMrJSZ6uEWc6v3yZo0ZhuCKq5c6D50r1b05Qp4SQ_iO2PQn0Ty8nFFIULWb_14ze9jlE8K4Aw9DkEr8VqQGpu8brWfnc_dK2SW2-yvEXk_lj61dOo6T49k2OcsZlU4wiBW9YZfzstiRghA70rySdzbcrdpKdYrJwd-FNNRtBGanU6ZgYv32Qp4oOzwpShtALI1mnL4A2-_Uc3BPWVvAewTQ_GAVNL5bkEHY_lU9kfq5XCeVAP4EzMrQyr7VYIqgpBml5GK8u8MZaDTy4AvdQmDIrZlxBltmVby-jFCajLRMGPNhHXc2Vp-AUbSSMp4w_ZdonQ_wTgaF1Uq2Xe4PqesSrKOuPLQMlun892NAVE-kz65fdcQWoRu9dYbRHIK7gpV271t5upYeg9D1VTaWPfHQ5toF9I4LQ9wCMNroj--nVusrafDwDtgpczZmjJNPbkZplbSMXC5UcLALmtpHN02SsyERM71c1oTUdlF_UP62mepXyP6NR8KbwjfeJUAkbNmpJm47OIqZ6jSy7EYTB2PPYdD_iVZrVH37g1LZTnHF0P3u2LgwnlZBHToS8Nj6uxamH437zMTCFYS7FKoP7OOS46isPbyCFiH-MaaNUAOngaLVMczSaF_yN43QXQs3Q9w7oZr5-BVWSxv3H3Qb_7WRoRcLsQJIyxLP9pptyCMmVmpq7bdywNRiNT-vQ5UObQss3T1AGiBavHWJFp3Cd0HZ8GzdOnwbl4rS9hL9V_8x36ldy_CLD0wUSLGVJMWuZlToJ-y5t03WqXJ8rAJRiSJKJdc9jtIuUfR7xhDeBnsMW3ftV7j13JXM2LZXP1V6ddLo0wgwSwOFkAL7Hm5Pruxc5iqhsrbCmOs0vKeI0cOlj2kcNkTc6dc0eMWZWAW6g9y6YdE0PxhRD9X7WH0dPH2WbZ6AsQWPzdwBdR5sx24n0-TmpLizC1Dguj_kyxDFpw1s0TfN1aBcMbSHF1oF9ECd9WbTdfiTnI9WOl_-XvUzWmETgLfESUQQAvRQqo1mqxLm5eiw9AOfZX-7qw6UBupDbhy0Rx9m3DR-C1tZQrnMn32o2XJ4H1-OFXGTP9P3fRIgF8m0jcfEEki0X6dCr-DS7e9uN7aLVaw8sOrCkgmogC0Zss5coZ5yNAIZFU24yxm45FD4N1NQz5s80zHZ-jXS4JUKNYaASR9Dd5WsBunlXeYrJndyRBxgvLunDpLArU6s0vJIUuYSsLbGmNTAf85T_-0U_b7XWoeW8R5y2fNzEaTTx3uBQImWCD6Y7wlr7dysJ-i1sw3QKb72J4UcwSUOqd-6yddplLY8g1vUTjLKFDwsF_C31hxHnTMb_tZ185_o3MV_y8GJkgz8FG9H4GXV4yioNTOL1GD9MOKNUk8lSiaQtoScKZs1_8LbVEd5yRy67tS8OtSWTtu61vwQfjwfb14D9wxBxlRAPdnkTegu2iSn5vI4Dtkke_DgShMI7U2pmr5Gk-7beLyOklhPNw2fPSZb-nafFI9Yulc8UvK8mG8bQ3wGi85dD7sd74xpJRXp7Oi3dMqZofh2Jo5aieIJqtE3C9Sv1DowC1beAPQLHHWonb32H1gJUPwZsRJ1stcumod0ws8Fngu1360I4SgWaC9mcDCF32Jwf2HOFkcrguYVdEAbZ9YwvAESksT2QNB-Lhjx9pfZCFG_Fz2Q-PWeKZUM_oOETNTJdNSrAG5rHmEurfm_hJkosjnrQ9SB7s6i3rYLmqZRu_QD29mFeo1Ute1XufKaVAwZthA5lXVxHK9cAbYGSE6bQ7lQ6ExiqhijKqe1PAPXq_wyIwTYy6_C7Fg-wZur3_U_ghGZDX5AFXDoGDByU5mO3vdd9Vj9_VOsEcR_hoEMlEA5FSXJ4aqV6-ztMXaPkGIzv6NF0gKqgKMfM0ObZttzvQxanE1wAl_9IFk3fV5RidL9FzFbNHQU6lqFfouDy10KNlSqCesg2yts6fL0CtatNjpnSt0DKknlljvPHK9ecPUCVprOvsMYNiWSsrmclAaZO96_1rb402bTwHvGh8pbzXHZagnk6p1uaOx4OsUf4_q6IErTgWOIpVByjlsv3br4jxGdiv9MWzNuX7K1lFgCE7uSflD69ZCiW8QwDgZLmw00nm4ELmraoqTtMzMIHaO2uehI1EnjzudftcQcbx8yN0ffno45RJr_b5D1OVNUBpM_2r5b8zgTZxwA9IpDw9TNuJauiqoiVWVr57vzy5MjPH9qzFNGo45qVf-F2qda-j9C0hPDNs7UP29auLaXGQGcMtF2Wjt6UWJaC3lEQ0lSyTwW5Mkhqllk7Cb988ZeY7N46chVjm_Orh92BRpj2zG8RFBcQPTLUH4qEjj8mhLCkXJYutd1sVSl3QV_Aw6GkkDwVqQMr2zTLCNmcvcDHmp7eG5GnEbtz6QS_oPc45TzO8YXkQTxQZ167hJ4PJuaiFM9Ub_Xtvqe6ZHA_IcMiJbF_t6jRCkiEuNvqfj-j5ZotHFhQxJLJc3jVvbGHKzb8DSMUB1CPAybvyBQcrxUnAtatVV42nN7UGpRBbNOz0ucDHwcYglupBaio6H9esUdqnL4JMsWxmbGp1bzwZEiFfN-xFbTxrkF23xnHXz_2lmCETLSudiyddvt3U2-jMEJnQH9Us'),
       ('Омерзительная восьмёрка', 'США после Гражданской войны. Легендарный охотник за головами Джон Рут по кличке Вешатель конвоирует заключенную. По пути к ним прибиваются еще несколько путешественников. Снежная буря вынуждает компанию искать укрытие в лавке на отшибе, где уже расположилось весьма пёстрое общество: генерал конфедератов, мексиканец, ковбой… И один из них — не тот, за кого себя выдает.', 'https://kinopoisk-ru.clstorage.net/2cN9w8350/dbcad21QyAj/ZLTywhm6vTuulVYDpQLu0uTGFm-McbyCgNpEscudfabVO--cl4oZvsCgjupsjzlCmjyakdyCnKmK9ZL2s0XzM_Dl8hXdPCtvuQZGFjPXLsK2RcO8nqOoU4eekKZcVrGs2bGzFRhAEbXY8R67qxhuFMo5zAeMFJo5GakLMm0QJaRq-bdHWaU7t6NEVgYlxiSWecMcq145s0rbW0p3DVO6p9hzfRpWD1A93TeS-jzS6JoFdozDjFOI9FBtDCPkUS1srjo-xlfpv3rjwp3AIspr3vWJlm_WtXnHGpWD8sIat-ZadPSalgFZtR9_1_652KJVAXENyERUA_fcJMrq5lMrJSZ6uEWc6v3yZo0ZhuCKq5c6D50r1b05Qp4SQ_iO2PQn0Ty8nFFIULWb_14ze9jlE8K4Aw9DkEr8VqQGpu8brWfnc_dK2SW2-yvEXk_lj61dOo6T49k2OcsZlU4wiBW9YZfzstiRghA70rySdzbcrdpKdYrJwd-FNNRtBGanU6ZgYv32Qp4oOzwpShtALI1mnL4A2-_Uc3BPWVvAewTQ_GAVNL5bkEHY_lU9kfq5XCeVAP4EzMrQyr7VYIqgpBml5GK8u8MZaDTy4AvdQmDIrZlxBltmVby-jFCajLRMGPNhHXc2Vp-AUbSSMp4w_ZdonQ_wTgaF1Uq2Xe4PqesSrKOuPLQMlun892NAVE-kz65fdcQWoRu9dYbRHIK7gpV271t5upYeg9D1VTaWPfHQ5toF9I4LQ9wCMNroj--nVusrafDwDtgpczZmjJNPbkZplbSMXC5UcLALmtpHN02SsyERM71c1oTUdlF_UP62mepXyP6NR8KbwjfeJUAkbNmpJm47OIqZ6jSy7EYTB2PPYdD_iVZrVH37g1LZTnHF0P3u2LgwnlZBHToS8Nj6uxamH437zMTCFYS7FKoP7OOS46isPbyCFiH-MaaNUAOngaLVMczSaF_yN43QXQs3Q9w7oZr5-BVWSxv3H3Qb_7WRoRcLsQJIyxLP9pptyCMmVmpq7bdywNRiNT-vQ5UObQss3T1AGiBavHWJFp3Cd0HZ8GzdOnwbl4rS9hL9V_8x36ldy_CLD0wUSLGVJMWuZlToJ-y5t03WqXJ8rAJRiSJKJdc9jtIuUfR7xhDeBnsMW3ftV7j13JXM2LZXP1V6ddLo0wgwSwOFkAL7Hm5Pruxc5iqhsrbCmOs0vKeI0cOlj2kcNkTc6dc0eMWZWAW6g9y6YdE0PxhRD9X7WH0dPH2WbZ6AsQWPzdwBdR5sx24n0-TmpLizC1Dguj_kyxDFpw1s0TfN1aBcMbSHF1oF9ECd9WbTdfiTnI9WOl_-XvUzWmETgLfESUQQAvRQqo1mqxLm5eiw9AOfZX-7qw6UBupDbhy0Rx9m3DR-C1tZQrnMn32o2XJ4H1-OFXGTP9P3fRIgF8m0jcfEEki0X6dCr-DS7e9uN7aLVaw8sOrCkgmogC0Zss5coZ5yNAIZFU24yxm45FD4N1NQz5s80zHZ-jXS4JUKNYaASR9Dd5WsBunlXeYrJndyRBxgvLunDpLArU6s0vJIUuYSsLbGmNTAf85T_-0U_b7XWoeW8R5y2fNzEaTTx3uBQImWCD6Y7wlr7dysJ-i1sw3QKb72J4UcwSUOqd-6yddplLY8g1vUTjLKFDwsF_C31hxHnTMb_tZ185_o3MV_y8GJkgz8FG9H4GXV4yioNTOL1GD9MOKNUk8lSiaQtoScKZs1_8LbVEd5yRy67tS8OtSWTtu61vwQfjwfb14D9wxBxlRAPdnkTegu2iSn5vI4Dtkke_DgShMI7U2pmr5Gk-7beLyOklhPNw2fPSZb-nafFI9Yulc8UvK8mG8bQ3wGi85dD7sd74xpJRXp7Oi3dMqZofh2Jo5aieIJqtE3C9Sv1DowC1beAPQLHHWonb32H1gJUPwZsRJ1stcumod0ws8Fngu1360I4SgWaC9mcDCF32Jwf2HOFkcrguYVdEAbZ9YwvAESksT2QNB-Lhjx9pfZCFG_Fz2Q-PWeKZUM_oOETNTJdNSrAG5rHmEurfm_hJkosjnrQ9SB7s6i3rYLmqZRu_QD29mFeo1Ute1XufKaVAwZthA5lXVxHK9cAbYGSE6bQ7lQ6ExiqhijKqe1PAPXq_wyIwTYy6_C7Fg-wZur3_U_ghGZDX5AFXDoGDByU5mO3vdd9Vj9_VOsEcR_hoEMlEA5FSXJ4aqV6-ztMXaPkGIzv6NF0gKqgKMfM0ObZttzvQxanE1wAl_9IFk3fV5RidL9FzFbNHQU6lqFfouDy10KNlSqCesg2yts6fL0CtatNjpnSt0DKknlljvPHK9ecPUCVprOvsMYNiWSsrmclAaZO96_1rb402bTwHvGh8pbzXHZagnk6p1uaOx4OsUf4_q6IErTgWOIpVByjlsv3br4jxGdiv9MWzNuX7K1lFgCE7uSflD69ZCiW8QwDgZLmw00nm4ELmraoqTtMzMIHaO2uehI1EnjzudftcQcbx8yN0ffno45RJr_b5D1OVNUBpM_2r5b8zgTZxwA9IpDw9TNuJauiqoiVWVr57vzy5MjPH9qzFNGo45qVf-F2qda-j9C0hPDNs7UP29auLaXGQGcMtF2Wjt6UWJaC3lEQ0lSyTwW5Mkhqllk7Cb988ZeY7N46chVjm_Orh92BRpj2zG8RFBcQPTLUH4qEjj8mhLCkXJYutd1sVSl3QV_Aw6GkkDwVqQMr2zTLCNmcvcDHmp7eG5GnEbtz6QS_oPc45TzO8YXkQTxQZ167hJ4PJuaiFM9Ub_Xtvqe6ZHA_IcMiJbF_t6jRCkiEuNvqfj-j5ZotHFhQxJLJc3jVvbGHKzb8DSMUB1CPAybvyBQcrxUnAtatVV42nN7UGpRBbNOz0ucDHwcYglupBaio6H9esUdqnL4JMsWxmbGp1bzwZEiFfN-xFbTxrkF23xnHXz_2lmCETLSudiyddvt3U2-jMEJnQH9Us'),
       ('Убить Билла', 'В беременную наёмную убийцу по кличке Чёрная Мамба во время бракосочетания стреляет человек по имени Билл. Но голова у женщины оказалась крепкой — пролежав четыре года в коме, бывшая невеста приходит в себя.', 'https://kinopoisk-ru.clstorage.net/2cN9w8350/dbcad21QyAj/ZLTywhm6vTuulVYDpQLu0uTGFm-McbyCgNpEscudfabVO--cl4oZvsCgjiguzrjBWj0bU8mXX73d9ZG3ZgUzM_En5tXLaKtuORLEV2eXehRjk9e_XiOoU4eekKZcVrGs2bGzFRhAEbXY8R67qxhuFMo5zAeMFJo5GakLMm0QJaRq-bdHWaU7t6NEVgYlxiSWecMcq145s0rbW0p3DVO6p9hzfRpWD1A93TeS-jzS6JoFdozDjFOI9FBtDCPkUS1srjo-xlfpv3rjwp3AIspr3vWJlm_WtXnHGpWD8sIat-ZadPSalgFZtR9_1_652KJVAXENyERUA_fcJMrq5lMrJSZ6uEWc6v3yZo0ZhuCKq5c6D50r1b05Qp4SQ_iO2PQn0Ty8nFFIULWb_14ze9jlE8K4Aw9DkEr8VqQGpu8brWfnc_dK2SW2-yvEXk_lj61dOo6T49k2OcsZlU4wiBW9YZfzstiRghA70rySdzbcrdpKdYrJwd-FNNRtBGanU6ZgYv32Qp4oOzwpShtALI1mnL4A2-_Uc3BPWVvAewTQ_GAVNL5bkEHY_lU9kfq5XCeVAP4EzMrQyr7VYIqgpBml5GK8u8MZaDTy4AvdQmDIrZlxBltmVby-jFCajLRMGPNhHXc2Vp-AUbSSMp4w_ZdonQ_wTgaF1Uq2Xe4PqesSrKOuPLQMlun892NAVE-kz65fdcQWoRu9dYbRHIK7gpV271t5upYeg9D1VTaWPfHQ5toF9I4LQ9wCMNroj--nVusrafDwDtgpczZmjJNPbkZplbSMXC5UcLALmtpHN02SsyERM71c1oTUdlF_UP62mepXyP6NR8KbwjfeJUAkbNmpJm47OIqZ6jSy7EYTB2PPYdD_iVZrVH37g1LZTnHF0P3u2LgwnlZBHToS8Nj6uxamH437zMTCFYS7FKoP7OOS46isPbyCFiH-MaaNUAOngaLVMczSaF_yN43QXQs3Q9w7oZr5-BVWSxv3H3Qb_7WRoRcLsQJIyxLP9pptyCMmVmpq7bdywNRiNT-vQ5UObQss3T1AGiBavHWJFp3Cd0HZ8GzdOnwbl4rS9hL9V_8x36ldy_CLD0wUSLGVJMWuZlToJ-y5t03WqXJ8rAJRiSJKJdc9jtIuUfR7xhDeBnsMW3ftV7j13JXM2LZXP1V6ddLo0wgwSwOFkAL7Hm5Pruxc5iqhsrbCmOs0vKeI0cOlj2kcNkTc6dc0eMWZWAW6g9y6YdE0PxhRD9X7WH0dPH2WbZ6AsQWPzdwBdR5sx24n0-TmpLizC1Dguj_kyxDFpw1s0TfN1aBcMbSHF1oF9ECd9WbTdfiTnI9WOl_-XvUzWmETgLfESUQQAvRQqo1mqxLm5eiw9AOfZX-7qw6UBupDbhy0Rx9m3DR-C1tZQrnMn32o2XJ4H1-OFXGTP9P3fRIgF8m0jcfEEki0X6dCr-DS7e9uN7aLVaw8sOrCkgmogC0Zss5coZ5yNAIZFU24yxm45FD4N1NQz5s80zHZ-jXS4JUKNYaASR9Dd5WsBunlXeYrJndyRBxgvLunDpLArU6s0vJIUuYSsLbGmNTAf85T_-0U_b7XWoeW8R5y2fNzEaTTx3uBQImWCD6Y7wlr7dysJ-i1sw3QKb72J4UcwSUOqd-6yddplLY8g1vUTjLKFDwsF_C31hxHnTMb_tZ185_o3MV_y8GJkgz8FG9H4GXV4yioNTOL1GD9MOKNUk8lSiaQtoScKZs1_8LbVEd5yRy67tS8OtSWTtu61vwQfjwfb14D9wxBxlRAPdnkTegu2iSn5vI4Dtkke_DgShMI7U2pmr5Gk-7beLyOklhPNw2fPSZb-nafFI9Yulc8UvK8mG8bQ3wGi85dD7sd74xpJRXp7Oi3dMqZofh2Jo5aieIJqtE3C9Sv1DowC1beAPQLHHWonb32H1gJUPwZsRJ1stcumod0ws8Fngu1360I4SgWaC9mcDCF32Jwf2HOFkcrguYVdEAbZ9YwvAESksT2QNB-Lhjx9pfZCFG_Fz2Q-PWeKZUM_oOETNTJdNSrAG5rHmEurfm_hJkosjnrQ9SB7s6i3rYLmqZRu_QD29mFeo1Ute1XufKaVAwZthA5lXVxHK9cAbYGSE6bQ7lQ6ExiqhijKqe1PAPXq_wyIwTYy6_C7Fg-wZur3_U_ghGZDX5AFXDoGDByU5mO3vdd9Vj9_VOsEcR_hoEMlEA5FSXJ4aqV6-ztMXaPkGIzv6NF0gKqgKMfM0ObZttzvQxanE1wAl_9IFk3fV5RidL9FzFbNHQU6lqFfouDy10KNlSqCesg2yts6fL0CtatNjpnSt0DKknlljvPHK9ecPUCVprOvsMYNiWSsrmclAaZO96_1rb402bTwHvGh8pbzXHZagnk6p1uaOx4OsUf4_q6IErTgWOIpVByjlsv3br4jxGdiv9MWzNuX7K1lFgCE7uSflD69ZCiW8QwDgZLmw00nm4ELmraoqTtMzMIHaO2uehI1EnjzudftcQcbx8yN0ffno45RJr_b5D1OVNUBpM_2r5b8zgTZxwA9IpDw9TNuJauiqoiVWVr57vzy5MjPH9qzFNGo45qVf-F2qda-j9C0hPDNs7UP29auLaXGQGcMtF2Wjt6UWJaC3lEQ0lSyTwW5Mkhqllk7Cb988ZeY7N46chVjm_Orh92BRpj2zG8RFBcQPTLUH4qEjj8mhLCkXJYutd1sVSl3QV_Aw6GkkDwVqQMr2zTLCNmcvcDHmp7eG5GnEbtz6QS_oPc45TzO8YXkQTxQZ167hJ4PJuaiFM9Ub_Xtvqe6ZHA_IcMiJbF_t6jRCkiEuNvqfj-j5ZotHFhQxJLJc3jVvbGHKzb8DSMUB1CPAybvyBQcrxUnAtatVV42nN7UGpRBbNOz0ucDHwcYglupBaio6H9esUdqnL4JMsWxmbGp1bzwZEiFfN-xFbTxrkF23xnHXz_2lmCETLSudiyddvt3U2-jMEJnQH9Us'),
       ('Крёстный отец 2', 'Для дона Корлеоне и его сына не существует моральных преград на пути к достижению целей. Они превращают мафию, построенную по патриархальным сицилийским законам, в весьма прагматичную, жесткую корпорацию, плавно интегрируя её в большой бизнес Америки.', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1599028/33474b2a-d670-47c8-9cbe-51291847b6d4/600x900'),
       ('Старикам тут не место', 'Обычный работяга обнаруживает в пустыне гору трупов, набитый героином грузовик и соблазнительную сумму в два миллиона долларов наличными. Он решает взять деньги себе, и результатом становится волна насилия, которую не может остановить вся полиция Западного Техаса.', 'https://kinopoisk-ru.clstorage.net/2cN9w8350/dbcad21QyAj/ZLTywhm6vTuulVYDpQLu0uTGFm-McbyCgNpEscudfabVO--cl4oZvsCgjiguzrjBWjxO0d2CXfwI9YSjZtFzM_Lwc5Xd6X9ueRKE17MDu0A2xUF_3-OoU4eekKZcVrGs2bGzFRhAEbXY8R67qxhuFMo5zAeMFJo5GakLMm0QJaRq-bdHWaU7t6NEVgYlxiSWecMcq145s0rbW0p3DVO6p9hzfRpWD1A93TeS-jzS6JoFdozDjFOI9FBtDCPkUS1srjo-xlfpv3rjwp3AIspr3vWJlm_WtXnHGpWD8sIat-ZadPSalgFZtR9_1_652KJVAXENyERUA_fcJMrq5lMrJSZ6uEWc6v3yZo0ZhuCKq5c6D50r1b05Qp4SQ_iO2PQn0Ty8nFFIULWb_14ze9jlE8K4Aw9DkEr8VqQGpu8brWfnc_dK2SW2-yvEXk_lj61dOo6T49k2OcsZlU4wiBW9YZfzstiRghA70rySdzbcrdpKdYrJwd-FNNRtBGanU6ZgYv32Qp4oOzwpShtALI1mnL4A2-_Uc3BPWVvAewTQ_GAVNL5bkEHY_lU9kfq5XCeVAP4EzMrQyr7VYIqgpBml5GK8u8MZaDTy4AvdQmDIrZlxBltmVby-jFCajLRMGPNhHXc2Vp-AUbSSMp4w_ZdonQ_wTgaF1Uq2Xe4PqesSrKOuPLQMlun892NAVE-kz65fdcQWoRu9dYbRHIK7gpV271t5upYeg9D1VTaWPfHQ5toF9I4LQ9wCMNroj--nVusrafDwDtgpczZmjJNPbkZplbSMXC5UcLALmtpHN02SsyERM71c1oTUdlF_UP62mepXyP6NR8KbwjfeJUAkbNmpJm47OIqZ6jSy7EYTB2PPYdD_iVZrVH37g1LZTnHF0P3u2LgwnlZBHToS8Nj6uxamH437zMTCFYS7FKoP7OOS46isPbyCFiH-MaaNUAOngaLVMczSaF_yN43QXQs3Q9w7oZr5-BVWSxv3H3Qb_7WRoRcLsQJIyxLP9pptyCMmVmpq7bdywNRiNT-vQ5UObQss3T1AGiBavHWJFp3Cd0HZ8GzdOnwbl4rS9hL9V_8x36ldy_CLD0wUSLGVJMWuZlToJ-y5t03WqXJ8rAJRiSJKJdc9jtIuUfR7xhDeBnsMW3ftV7j13JXM2LZXP1V6ddLo0wgwSwOFkAL7Hm5Pruxc5iqhsrbCmOs0vKeI0cOlj2kcNkTc6dc0eMWZWAW6g9y6YdE0PxhRD9X7WH0dPH2WbZ6AsQWPzdwBdR5sx24n0-TmpLizC1Dguj_kyxDFpw1s0TfN1aBcMbSHF1oF9ECd9WbTdfiTnI9WOl_-XvUzWmETgLfESUQQAvRQqo1mqxLm5eiw9AOfZX-7qw6UBupDbhy0Rx9m3DR-C1tZQrnMn32o2XJ4H1-OFXGTP9P3fRIgF8m0jcfEEki0X6dCr-DS7e9uN7aLVaw8sOrCkgmogC0Zss5coZ5yNAIZFU24yxm45FD4N1NQz5s80zHZ-jXS4JUKNYaASR9Dd5WsBunlXeYrJndyRBxgvLunDpLArU6s0vJIUuYSsLbGmNTAf85T_-0U_b7XWoeW8R5y2fNzEaTTx3uBQImWCD6Y7wlr7dysJ-i1sw3QKb72J4UcwSUOqd-6yddplLY8g1vUTjLKFDwsF_C31hxHnTMb_tZ185_o3MV_y8GJkgz8FG9H4GXV4yioNTOL1GD9MOKNUk8lSiaQtoScKZs1_8LbVEd5yRy67tS8OtSWTtu61vwQfjwfb14D9wxBxlRAPdnkTegu2iSn5vI4Dtkke_DgShMI7U2pmr5Gk-7beLyOklhPNw2fPSZb-nafFI9Yulc8UvK8mG8bQ3wGi85dD7sd74xpJRXp7Oi3dMqZofh2Jo5aieIJqtE3C9Sv1DowC1beAPQLHHWonb32H1gJUPwZsRJ1stcumod0ws8Fngu1360I4SgWaC9mcDCF32Jwf2HOFkcrguYVdEAbZ9YwvAESksT2QNB-Lhjx9pfZCFG_Fz2Q-PWeKZUM_oOETNTJdNSrAG5rHmEurfm_hJkosjnrQ9SB7s6i3rYLmqZRu_QD29mFeo1Ute1XufKaVAwZthA5lXVxHK9cAbYGSE6bQ7lQ6ExiqhijKqe1PAPXq_wyIwTYy6_C7Fg-wZur3_U_ghGZDX5AFXDoGDByU5mO3vdd9Vj9_VOsEcR_hoEMlEA5FSXJ4aqV6-ztMXaPkGIzv6NF0gKqgKMfM0ObZttzvQxanE1wAl_9IFk3fV5RidL9FzFbNHQU6lqFfouDy10KNlSqCesg2yts6fL0CtatNjpnSt0DKknlljvPHK9ecPUCVprOvsMYNiWSsrmclAaZO96_1rb402bTwHvGh8pbzXHZagnk6p1uaOx4OsUf4_q6IErTgWOIpVByjlsv3br4jxGdiv9MWzNuX7K1lFgCE7uSflD69ZCiW8QwDgZLmw00nm4ELmraoqTtMzMIHaO2uehI1EnjzudftcQcbx8yN0ffno45RJr_b5D1OVNUBpM_2r5b8zgTZxwA9IpDw9TNuJauiqoiVWVr57vzy5MjPH9qzFNGo45qVf-F2qda-j9C0hPDNs7UP29auLaXGQGcMtF2Wjt6UWJaC3lEQ0lSyTwW5Mkhqllk7Cb988ZeY7N46chVjm_Orh92BRpj2zG8RFBcQPTLUH4qEjj8mhLCkXJYutd1sVSl3QV_Aw6GkkDwVqQMr2zTLCNmcvcDHmp7eG5GnEbtz6QS_oPc45TzO8YXkQTxQZ167hJ4PJuaiFM9Ub_Xtvqe6ZHA_IcMiJbF_t6jRCkiEuNvqfj-j5ZotHFhQxJLJc3jVvbGHKzb8DSMUB1CPAybvyBQcrxUnAtatVV42nN7UGpRBbNOz0ucDHwcYglupBaio6H9esUdqnL4JMsWxmbGp1bzwZEiFfN-xFbTxrkF23xnHXz_2lmCETLSudiyddvt3U2-jMEJnQH9Us'),
       ('Крёстный отец', 'Криминальная сага, повествующая о нью-йоркской сицилийской мафиозной семье Корлеоне. Фильм охватывает период 1945-1955 годов.', 'https://kinopoisk-ru.clstorage.net/2cN9w8350/dbcad21QyAj/ZLTywhm6vTuulVYDpQLu0uTGFm-McbyCgNpEscudfabVO--cl4oZvsCgjiguzrjBWj0bkZ1DHTzK9ZF3MlBzM_ExJtXLaGuuORKFFrIXOcK3EEL_3yOoU4eekKZcVrGs2bGzFRhAEbXY8R67qxhuFMo5zAeMFJo5GakLMm0QJaRq-bdHWaU7t6NEVgYlxiSWecMcq145s0rbW0p3DVO6p9hzfRpWD1A93TeS-jzS6JoFdozDjFOI9FBtDCPkUS1srjo-xlfpv3rjwp3AIspr3vWJlm_WtXnHGpWD8sIat-ZadPSalgFZtR9_1_652KJVAXENyERUA_fcJMrq5lMrJSZ6uEWc6v3yZo0ZhuCKq5c6D50r1b05Qp4SQ_iO2PQn0Ty8nFFIULWb_14ze9jlE8K4Aw9DkEr8VqQGpu8brWfnc_dK2SW2-yvEXk_lj61dOo6T49k2OcsZlU4wiBW9YZfzstiRghA70rySdzbcrdpKdYrJwd-FNNRtBGanU6ZgYv32Qp4oOzwpShtALI1mnL4A2-_Uc3BPWVvAewTQ_GAVNL5bkEHY_lU9kfq5XCeVAP4EzMrQyr7VYIqgpBml5GK8u8MZaDTy4AvdQmDIrZlxBltmVby-jFCajLRMGPNhHXc2Vp-AUbSSMp4w_ZdonQ_wTgaF1Uq2Xe4PqesSrKOuPLQMlun892NAVE-kz65fdcQWoRu9dYbRHIK7gpV271t5upYeg9D1VTaWPfHQ5toF9I4LQ9wCMNroj--nVusrafDwDtgpczZmjJNPbkZplbSMXC5UcLALmtpHN02SsyERM71c1oTUdlF_UP62mepXyP6NR8KbwjfeJUAkbNmpJm47OIqZ6jSy7EYTB2PPYdD_iVZrVH37g1LZTnHF0P3u2LgwnlZBHToS8Nj6uxamH437zMTCFYS7FKoP7OOS46isPbyCFiH-MaaNUAOngaLVMczSaF_yN43QXQs3Q9w7oZr5-BVWSxv3H3Qb_7WRoRcLsQJIyxLP9pptyCMmVmpq7bdywNRiNT-vQ5UObQss3T1AGiBavHWJFp3Cd0HZ8GzdOnwbl4rS9hL9V_8x36ldy_CLD0wUSLGVJMWuZlToJ-y5t03WqXJ8rAJRiSJKJdc9jtIuUfR7xhDeBnsMW3ftV7j13JXM2LZXP1V6ddLo0wgwSwOFkAL7Hm5Pruxc5iqhsrbCmOs0vKeI0cOlj2kcNkTc6dc0eMWZWAW6g9y6YdE0PxhRD9X7WH0dPH2WbZ6AsQWPzdwBdR5sx24n0-TmpLizC1Dguj_kyxDFpw1s0TfN1aBcMbSHF1oF9ECd9WbTdfiTnI9WOl_-XvUzWmETgLfESUQQAvRQqo1mqxLm5eiw9AOfZX-7qw6UBupDbhy0Rx9m3DR-C1tZQrnMn32o2XJ4H1-OFXGTP9P3fRIgF8m0jcfEEki0X6dCr-DS7e9uN7aLVaw8sOrCkgmogC0Zss5coZ5yNAIZFU24yxm45FD4N1NQz5s80zHZ-jXS4JUKNYaASR9Dd5WsBunlXeYrJndyRBxgvLunDpLArU6s0vJIUuYSsLbGmNTAf85T_-0U_b7XWoeW8R5y2fNzEaTTx3uBQImWCD6Y7wlr7dysJ-i1sw3QKb72J4UcwSUOqd-6yddplLY8g1vUTjLKFDwsF_C31hxHnTMb_tZ185_o3MV_y8GJkgz8FG9H4GXV4yioNTOL1GD9MOKNUk8lSiaQtoScKZs1_8LbVEd5yRy67tS8OtSWTtu61vwQfjwfb14D9wxBxlRAPdnkTegu2iSn5vI4Dtkke_DgShMI7U2pmr5Gk-7beLyOklhPNw2fPSZb-nafFI9Yulc8UvK8mG8bQ3wGi85dD7sd74xpJRXp7Oi3dMqZofh2Jo5aieIJqtE3C9Sv1DowC1beAPQLHHWonb32H1gJUPwZsRJ1stcumod0ws8Fngu1360I4SgWaC9mcDCF32Jwf2HOFkcrguYVdEAbZ9YwvAESksT2QNB-Lhjx9pfZCFG_Fz2Q-PWeKZUM_oOETNTJdNSrAG5rHmEurfm_hJkosjnrQ9SB7s6i3rYLmqZRu_QD29mFeo1Ute1XufKaVAwZthA5lXVxHK9cAbYGSE6bQ7lQ6ExiqhijKqe1PAPXq_wyIwTYy6_C7Fg-wZur3_U_ghGZDX5AFXDoGDByU5mO3vdd9Vj9_VOsEcR_hoEMlEA5FSXJ4aqV6-ztMXaPkGIzv6NF0gKqgKMfM0ObZttzvQxanE1wAl_9IFk3fV5RidL9FzFbNHQU6lqFfouDy10KNlSqCesg2yts6fL0CtatNjpnSt0DKknlljvPHK9ecPUCVprOvsMYNiWSsrmclAaZO96_1rb402bTwHvGh8pbzXHZagnk6p1uaOx4OsUf4_q6IErTgWOIpVByjlsv3br4jxGdiv9MWzNuX7K1lFgCE7uSflD69ZCiW8QwDgZLmw00nm4ELmraoqTtMzMIHaO2uehI1EnjzudftcQcbx8yN0ffno45RJr_b5D1OVNUBpM_2r5b8zgTZxwA9IpDw9TNuJauiqoiVWVr57vzy5MjPH9qzFNGo45qVf-F2qda-j9C0hPDNs7UP29auLaXGQGcMtF2Wjt6UWJaC3lEQ0lSyTwW5Mkhqllk7Cb988ZeY7N46chVjm_Orh92BRpj2zG8RFBcQPTLUH4qEjj8mhLCkXJYutd1sVSl3QV_Aw6GkkDwVqQMr2zTLCNmcvcDHmp7eG5GnEbtz6QS_oPc45TzO8YXkQTxQZ167hJ4PJuaiFM9Ub_Xtvqe6ZHA_IcMiJbF_t6jRCkiEuNvqfj-j5ZotHFhQxJLJc3jVvbGHKzb8DSMUB1CPAybvyBQcrxUnAtatVV42nN7UGpRBbNOz0ucDHwcYglupBaio6H9esUdqnL4JMsWxmbGp1bzwZEiFfN-xFbTxrkF23xnHXz_2lmCETLSudiyddvt3U2-jMEJnQH9Us');

INSERT INTO comments
(content, user_id, movie_id)
VALUES ('комментарий 1', 1, 1),
       ('комментарий 2', 2, 1);