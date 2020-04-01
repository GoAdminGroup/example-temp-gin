--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.16
-- Dumped by pg_dump version 12.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.goadmin_menu_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_menu_myid_seq OWNER TO golang;

SET default_tablespace = '';

--
-- Name: goadmin_menu; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_menu (
    id integer DEFAULT nextval('public.goadmin_menu_myid_seq'::regclass) NOT NULL,
    parent_id integer DEFAULT 0 NOT NULL,
    type integer DEFAULT 0,
    "order" integer DEFAULT 0 NOT NULL,
    title character varying(50) NOT NULL,
    header character varying(100),
    icon character varying(50) NOT NULL,
    uri character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_menu OWNER TO golang;

--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.goadmin_operation_log_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_operation_log_myid_seq OWNER TO golang;

--
-- Name: goadmin_operation_log; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_operation_log (
    id integer DEFAULT nextval('public.goadmin_operation_log_myid_seq'::regclass) NOT NULL,
    user_id integer NOT NULL,
    path character varying(255) NOT NULL,
    method character varying(10) NOT NULL,
    ip character varying(15) NOT NULL,
    input text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_operation_log OWNER TO golang;

--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.goadmin_permissions_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_permissions_myid_seq OWNER TO golang;

--
-- Name: goadmin_permissions; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_permissions (
    id integer DEFAULT nextval('public.goadmin_permissions_myid_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    slug character varying(50) NOT NULL,
    http_method character varying(255),
    http_path text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_permissions OWNER TO golang;

--
-- Name: goadmin_role_menu; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_role_menu (
    role_id integer NOT NULL,
    menu_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_menu OWNER TO golang;

--
-- Name: goadmin_role_permissions; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_role_permissions (
    role_id integer NOT NULL,
    permission_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_permissions OWNER TO golang;

--
-- Name: goadmin_role_users; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_role_users (
    role_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_role_users OWNER TO golang;

--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.goadmin_roles_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_roles_myid_seq OWNER TO golang;

--
-- Name: goadmin_roles; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_roles (
    id integer DEFAULT nextval('public.goadmin_roles_myid_seq'::regclass) NOT NULL,
    name character varying NOT NULL,
    slug character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_roles OWNER TO golang;

--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.goadmin_session_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_session_myid_seq OWNER TO golang;

--
-- Name: goadmin_session; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_session (
    id integer DEFAULT nextval('public.goadmin_session_myid_seq'::regclass) NOT NULL,
    sid character varying(50) NOT NULL,
    "values" character varying(3000) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_session OWNER TO golang;

--
-- Name: goadmin_user_permissions; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_user_permissions (
    user_id integer NOT NULL,
    permission_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_user_permissions OWNER TO golang;

--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE; Schema: public; Owner: golang
--

CREATE SEQUENCE public.goadmin_users_myid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999
    CACHE 1;


ALTER TABLE public.goadmin_users_myid_seq OWNER TO golang;

--
-- Name: goadmin_users; Type: TABLE; Schema: public; Owner: golang
--

CREATE TABLE public.goadmin_users (
    id integer DEFAULT nextval('public.goadmin_users_myid_seq'::regclass) NOT NULL,
    username character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    avatar character varying(255),
    remember_token character varying(100),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goadmin_users OWNER TO golang;

--
-- Data for Name: goadmin_menu; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_menu (id, parent_id, type, "order", title, header, icon, uri, created_at, updated_at) FROM stdin;
7	0	1	1	Dashboard	\N	fa-bar-chart	/	2019-09-10 00:00:00	2019-09-10 00:00:00
8	0	0	2	Demo	Demo	fa-align-left		2020-02-16 11:54:06.601098	2020-02-16 11:54:06.601098
9	8	0	2	user_info	user_info	fa-bars	/info/user_info	2020-02-16 11:54:29.921597	2020-02-16 11:54:29.921597
10	8	0	3	user_rule	user_rule	fa-bars	/info/user_rule	2020-02-16 11:54:41.091392	2020-02-16 11:54:41.091392
11	8	0	4	user_info_rule	user_info_rule	fa-bars	/info/user_info_rule	2020-02-16 11:54:55.614	2020-02-16 11:54:55.614
1	0	1	5	Admin	\N	fa-tasks		2019-09-10 00:00:00	2019-09-10 00:00:00
2	1	1	5	Users	\N	fa-users	/info/manager	2019-09-10 00:00:00	2019-09-10 00:00:00
3	1	1	6	Roles	\N	fa-user	/info/roles	2019-09-10 00:00:00	2019-09-10 00:00:00
4	1	1	7	Permission	\N	fa-ban	/info/permission	2019-09-10 00:00:00	2019-09-10 00:00:00
5	1	1	8	Menu	\N	fa-bars	/menu	2019-09-10 00:00:00	2019-09-10 00:00:00
6	1	1	9	Operation log	\N	fa-history	/info/op	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Data for Name: goadmin_operation_log; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_operation_log (id, user_id, path, method, ip, input, created_at, updated_at) FROM stdin;
2	1	/admin/info/permission	GET	192.168.110.219		2020-02-16 11:53:01.311069	2020-02-16 11:53:01.311069
3	1	/admin/info/roles	GET	192.168.110.219		2020-02-16 11:53:06.027963	2020-02-16 11:53:06.027963
4	1	/admin/menu	GET	192.168.110.219		2020-02-16 11:53:09.513856	2020-02-16 11:53:09.513856
5	1	/admin/menu/new	POST	192.168.110.219	{"_previous_":["/admin/menu"],"_t":["1e516d96-4ae1-4caa-a5b0-b2dc8c4ba405"],"header":["Demo"],"icon":["fa-align-left"],"parent_id":["0"],"roles[]":["1","2"],"title":["Demo"],"uri":[""]}	2020-02-16 11:54:06.71943	2020-02-16 11:54:06.71943
6	1	/admin/menu/new	POST	192.168.110.219	{"_previous_":["/admin/menu"],"_t":["a38d35e8-6f46-47fd-9b10-848d9b508b67"],"header":["user_info"],"icon":["fa-bars"],"parent_id":["8"],"roles[]":["1","2"],"title":["user_info"],"uri":["/info/user_info"]}	2020-02-16 11:54:30.043361	2020-02-16 11:54:30.043361
7	1	/admin/menu/new	POST	192.168.110.219	{"_previous_":["/admin/menu"],"_t":["90cd48c3-e718-4c65-ad1c-7eee5984423a"],"header":["user_rule"],"icon":["fa-bars"],"parent_id":["8"],"title":["user_rule"],"uri":["/info/user_rule"]}	2020-02-16 11:54:41.199638	2020-02-16 11:54:41.199638
8	1	/admin/menu/new	POST	192.168.110.219	{"_previous_":["/admin/menu"],"_t":["ebe8540d-7594-4814-973e-490601ec2de0"],"header":["user_info_rule"],"icon":["fa-bars"],"parent_id":["8"],"roles[]":["1","2"],"title":["user_info_rule"],"uri":["/info/user_info_rule"]}	2020-02-16 11:54:55.748195	2020-02-16 11:54:55.748195
9	1	/admin/menu	GET	192.168.110.219		2020-02-16 11:55:00.828149	2020-02-16 11:55:00.828149
10	1	/admin/menu	GET	192.168.110.219		2020-02-16 11:55:05.704639	2020-02-16 11:55:05.704639
11	1	/admin/menu/order	POST	192.168.110.219		2020-02-16 11:55:09.576448	2020-02-16 11:55:09.576448
12	1	/admin/menu	GET	192.168.110.219		2020-02-16 11:55:09.693195	2020-02-16 11:55:09.693195
13	1	/admin/menu	GET	192.168.110.219		2020-02-16 11:55:10.789727	2020-02-16 11:55:10.789727
14	1	/admin/info/permission	GET	192.168.110.219		2020-02-16 11:55:14.248375	2020-02-16 11:55:14.248375
15	1	/admin/info/permission/new	GET	192.168.110.219		2020-02-16 11:55:24.943085	2020-02-16 11:55:24.943085
16	1	/admin/new/permission	POST	192.168.110.219	{"__go_admin_post_type":["1"],"_previous_":["/admin/info/permission?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["70346049-9eb5-4c49-a95b-b2c7b95c2d1a"],"http_method[]":[""],"http_path":["/info/user_info\\r\\n/info/user_info/new\\r\\n/info/user_info/edit\\r\\n/new/user_info\\r\\n/edit/user_info\\r\\n/export/user_info"],"id":["3"],"name":["Demo-user_info-*"],"slug":["Demo-user_info-*"]}	2020-02-16 11:56:35.879754	2020-02-16 11:56:35.879754
17	1	/admin/info/permission/new	GET	192.168.110.219		2020-02-16 11:56:37.964466	2020-02-16 11:56:37.964466
18	1	/admin/new/permission	POST	192.168.110.219	{"__go_admin_post_type":["1"],"_previous_":["/admin/info/permission?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["918d1164-ef26-456d-9ca3-e8948ee3371a"],"http_method[]":[""],"http_path":["/info/user_rule\\r\\n/info/user_rule/new\\r\\n/info/user_rule/edit\\r\\n/new/user_rule\\r\\n/edit/user_rule\\r\\n/export/user_rule"],"id":["4"],"name":["Demo-user_rule-*"],"slug":["Demo-user_rule-*"]}	2020-02-16 11:57:10.450322	2020-02-16 11:57:10.450322
19	1	/admin/info/permission/new	GET	192.168.110.219		2020-02-16 11:57:12.242211	2020-02-16 11:57:12.242211
20	1	/admin/new/permission	POST	192.168.110.219	{"__go_admin_post_type":["1"],"_previous_":["/admin/info/permission?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["10dab61d-a202-44ae-8d88-42ae68878c8c"],"http_method[]":[""],"http_path":["/info/user_info_rule\\r\\n/info/user_info_rule/new\\r\\n/info/user_info_rule/edit\\r\\n/new/user_info_rule\\r\\n/edit/user_info_rule\\r\\n/export/user_info_rule"],"id":["5"],"name":["Demo-user_info_rule-*"],"slug":["Demo-user_info_rule-*"]}	2020-02-16 11:57:32.937796	2020-02-16 11:57:32.937796
21	1	/admin/info/roles	GET	192.168.110.219		2020-02-16 11:57:36.342776	2020-02-16 11:57:36.342776
22	1	/admin/info/roles/edit	GET	192.168.110.219		2020-02-16 11:57:38.959591	2020-02-16 11:57:38.959591
23	1	/admin/edit/roles	POST	192.168.110.219	{"_previous_":["/admin/info/roles?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc"],"_t":["ec7041b3-94f0-4c84-b5e7-3dc7534863dd"],"id":["1"],"name":["Administrator"],"permission_id[]":["1","2","3","4","5",""],"slug":["administrator"]}	2020-02-16 11:57:43.160813	2020-02-16 11:57:43.160813
24	1	/admin/info/roles	GET	192.168.110.219		2020-02-16 11:57:50.225575	2020-02-16 11:57:50.225575
25	1	/admin/info/roles	GET	192.168.110.219		2020-02-16 11:57:54.538156	2020-02-16 11:57:54.538156
26	1	/admin/info/user_info	GET	192.168.110.219		2020-02-16 11:57:58.319909	2020-02-16 11:57:58.319909
27	1	/admin/info/user_rule	GET	192.168.110.219		2020-02-16 11:58:00.465634	2020-02-16 11:58:00.465634
28	1	/admin/info/user_info	GET	192.168.110.219		2020-02-16 11:58:01.409256	2020-02-16 11:58:01.409256
29	1	/admin/info/user_rule	GET	192.168.110.219		2020-02-16 11:58:02.251245	2020-02-16 11:58:02.251245
30	1	/admin/info/user_info_rule	GET	192.168.110.219		2020-02-16 11:58:02.877462	2020-02-16 11:58:02.877462
31	1	/admin/info/user_info	GET	192.168.110.219		2020-02-16 11:58:04.020895	2020-02-16 11:58:04.020895
32	1	/admin/info/user_info	GET	192.168.110.219		2020-02-16 11:58:49.772931	2020-02-16 11:58:49.772931
33	1	/admin/info/user_rule	GET	192.168.110.219		2020-02-16 11:58:51.282022	2020-02-16 11:58:51.282022
34	1	/admin/info/user_info	GET	192.168.110.219		2020-02-16 11:58:52.192289	2020-02-16 11:58:52.192289
35	1	/admin/info/user_info_rule	GET	192.168.110.219		2020-02-16 11:58:53.110981	2020-02-16 11:58:53.110981
36	1	/admin/info/user_rule	GET	192.168.110.219		2020-02-16 11:58:54.492993	2020-02-16 11:58:54.492993
37	1	/admin/info/user_info_rule	GET	192.168.110.219		2020-02-16 11:58:55.281834	2020-02-16 11:58:55.281834
38	1	/admin/info/user_rule	GET	192.168.110.219		2020-02-16 11:59:00.411366	2020-02-16 11:59:00.411366
39	1	/admin/info/user_info_rule	GET	192.168.110.219		2020-02-16 11:59:01.007953	2020-02-16 11:59:01.007953
40	1	/admin/info/user_rule	GET	192.168.110.219		2020-02-16 11:59:05.424971	2020-02-16 11:59:05.424971
41	1	/admin/info/user_info	GET	192.168.110.219		2020-02-16 11:59:18.860209	2020-02-16 11:59:18.860209
42	1	/admin/info/user_rule	GET	192.168.110.219		2020-02-16 11:59:19.494689	2020-02-16 11:59:19.494689
\.


--
-- Data for Name: goadmin_permissions; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_permissions (id, name, slug, http_method, http_path, created_at, updated_at) FROM stdin;
1	All permission	*		*	2019-09-10 00:00:00	2019-09-10 00:00:00
2	Dashboard	dashboard	GET,PUT,POST,DELETE	/	2019-09-10 00:00:00	2019-09-10 00:00:00
3	Demo-user_info-*	Demo-user_info-*		/info/user_info\r\n/info/user_info/new\r\n/info/user_info/edit\r\n/new/user_info\r\n/edit/user_info\r\n/export/user_info	2020-02-16 11:56:35.698992	2020-02-16 19:56:35
4	Demo-user_rule-*	Demo-user_rule-*		/info/user_rule\r\n/info/user_rule/new\r\n/info/user_rule/edit\r\n/new/user_rule\r\n/edit/user_rule\r\n/export/user_rule	2020-02-16 11:57:10.267172	2020-02-16 19:57:10
5	Demo-user_info_rule-*	Demo-user_info_rule-*		/info/user_info_rule\r\n/info/user_info_rule/new\r\n/info/user_info_rule/edit\r\n/new/user_info_rule\r\n/edit/user_info_rule\r\n/export/user_info_rule	2020-02-16 11:57:32.759763	2020-02-16 19:57:32
\.


--
-- Data for Name: goadmin_role_menu; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_role_menu (role_id, menu_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
1	7	2019-09-10 00:00:00	2019-09-10 00:00:00
2	7	2019-09-10 00:00:00	2019-09-10 00:00:00
1	8	2020-02-16 11:54:06.611421	2020-02-16 11:54:06.611421
2	8	2020-02-16 11:54:06.619331	2020-02-16 11:54:06.619331
1	9	2020-02-16 11:54:29.930539	2020-02-16 11:54:29.930539
2	9	2020-02-16 11:54:29.938488	2020-02-16 11:54:29.938488
1	11	2020-02-16 11:54:55.624628	2020-02-16 11:54:55.624628
2	11	2020-02-16 11:54:55.634709	2020-02-16 11:54:55.634709
\.


--
-- Data for Name: goadmin_role_permissions; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_role_permissions (role_id, permission_id, created_at, updated_at) FROM stdin;
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
0	3	\N	\N
1	1	2020-02-16 11:57:43.055442	2020-02-16 11:57:43.055442
1	2	2020-02-16 11:57:43.062245	2020-02-16 11:57:43.062245
1	3	2020-02-16 11:57:43.066774	2020-02-16 11:57:43.066774
1	4	2020-02-16 11:57:43.072757	2020-02-16 11:57:43.072757
1	5	2020-02-16 11:57:43.077505	2020-02-16 11:57:43.077505
\.


--
-- Data for Name: goadmin_role_users; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_role_users (role_id, user_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Data for Name: goadmin_roles; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_roles (id, name, slug, created_at, updated_at) FROM stdin;
2	Operator	operator	2019-09-10 00:00:00	2019-09-10 00:00:00
1	Administrator	administrator	2019-09-10 00:00:00	2020-02-16 19:57:43
\.


--
-- Data for Name: goadmin_session; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_session (id, sid, "values", created_at, updated_at) FROM stdin;
2	7421c675-6da1-43a3-8380-2735a9419d8a	{"user_id":1}	2020-02-16 11:52:56.996287	2020-02-16 11:52:56.996287
\.


--
-- Data for Name: goadmin_user_permissions; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_user_permissions (user_id, permission_id, created_at, updated_at) FROM stdin;
1	1	2019-09-10 00:00:00	2019-09-10 00:00:00
2	2	2019-09-10 00:00:00	2019-09-10 00:00:00
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
0	1	\N	\N
\.


--
-- Data for Name: goadmin_users; Type: TABLE DATA; Schema: public; Owner: golang
--

COPY public.goadmin_users (id, username, password, name, avatar, remember_token, created_at, updated_at) FROM stdin;
2	operator	$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.	Operator		\N	2019-09-10 00:00:00	2019-09-10 00:00:00
1	admin	$2a$10$BO/dMCEhIpUKj.mC/LlRP.P1FnOfjDviacJzYEZK8HAslKFAiweui	admin		tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh	2019-09-10 00:00:00	2019-09-10 00:00:00
\.


--
-- Name: goadmin_menu_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: golang
--

SELECT pg_catalog.setval('public.goadmin_menu_myid_seq', 11, true);


--
-- Name: goadmin_operation_log_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: golang
--

SELECT pg_catalog.setval('public.goadmin_operation_log_myid_seq', 42, true);


--
-- Name: goadmin_permissions_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: golang
--

SELECT pg_catalog.setval('public.goadmin_permissions_myid_seq', 5, true);


--
-- Name: goadmin_roles_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: golang
--

SELECT pg_catalog.setval('public.goadmin_roles_myid_seq', 2, true);


--
-- Name: goadmin_session_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: golang
--

SELECT pg_catalog.setval('public.goadmin_session_myid_seq', 2, true);


--
-- Name: goadmin_users_myid_seq; Type: SEQUENCE SET; Schema: public; Owner: golang
--

SELECT pg_catalog.setval('public.goadmin_users_myid_seq', 2, true);


--
-- Name: goadmin_menu goadmin_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.goadmin_menu
    ADD CONSTRAINT goadmin_menu_pkey PRIMARY KEY (id);


--
-- Name: goadmin_operation_log goadmin_operation_log_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.goadmin_operation_log
    ADD CONSTRAINT goadmin_operation_log_pkey PRIMARY KEY (id);


--
-- Name: goadmin_permissions goadmin_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.goadmin_permissions
    ADD CONSTRAINT goadmin_permissions_pkey PRIMARY KEY (id);


--
-- Name: goadmin_roles goadmin_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.goadmin_roles
    ADD CONSTRAINT goadmin_roles_pkey PRIMARY KEY (id);


--
-- Name: goadmin_session goadmin_session_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.goadmin_session
    ADD CONSTRAINT goadmin_session_pkey PRIMARY KEY (id);


--
-- Name: goadmin_users goadmin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: golang
--

ALTER TABLE ONLY public.goadmin_users
    ADD CONSTRAINT goadmin_users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

