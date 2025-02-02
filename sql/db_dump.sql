--
-- PostgreSQL database dump
--

-- Dumped from database version 14.13
-- Dumped by pg_dump version 15.6

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
-- Name: btree_gin; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS btree_gin WITH SCHEMA public;


--
-- Name: EXTENSION btree_gin; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION btree_gin IS 'support for indexing common datatypes in GIN';


--
-- Name: btree_gist; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS btree_gist WITH SCHEMA public;


--
-- Name: EXTENSION btree_gist; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION btree_gist IS 'support for indexing common datatypes in GiST';


--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- Name: cube; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS cube WITH SCHEMA public;


--
-- Name: EXTENSION cube; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION cube IS 'data type for multidimensional cubes';


--
-- Name: dblink; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS dblink WITH SCHEMA public;


--
-- Name: EXTENSION dblink; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION dblink IS 'connect to other PostgreSQL databases from within a database';


--
-- Name: dict_int; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS dict_int WITH SCHEMA public;


--
-- Name: EXTENSION dict_int; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION dict_int IS 'text search dictionary template for integers';


--
-- Name: dict_xsyn; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS dict_xsyn WITH SCHEMA public;


--
-- Name: EXTENSION dict_xsyn; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION dict_xsyn IS 'text search dictionary template for extended synonym processing';


--
-- Name: earthdistance; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS earthdistance WITH SCHEMA public;


--
-- Name: EXTENSION earthdistance; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION earthdistance IS 'calculate great-circle distances on the surface of the Earth';


--
-- Name: fuzzystrmatch; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS fuzzystrmatch WITH SCHEMA public;


--
-- Name: EXTENSION fuzzystrmatch; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION fuzzystrmatch IS 'determine similarities and distance between strings';


--
-- Name: hstore; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS hstore WITH SCHEMA public;


--
-- Name: EXTENSION hstore; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION hstore IS 'data type for storing sets of (key, value) pairs';


--
-- Name: intarray; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS intarray WITH SCHEMA public;


--
-- Name: EXTENSION intarray; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION intarray IS 'functions, operators, and index support for 1-D arrays of integers';


--
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


--
-- Name: pg_stat_statements; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA public;


--
-- Name: EXTENSION pg_stat_statements; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_stat_statements IS 'track execution statistics of all SQL statements executed';


--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: pgrowlocks; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgrowlocks WITH SCHEMA public;


--
-- Name: EXTENSION pgrowlocks; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgrowlocks IS 'show row-level locking information';


--
-- Name: pgstattuple; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgstattuple WITH SCHEMA public;


--
-- Name: EXTENSION pgstattuple; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgstattuple IS 'show tuple-level statistics';


--
-- Name: tablefunc; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS tablefunc WITH SCHEMA public;


--
-- Name: EXTENSION tablefunc; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION tablefunc IS 'functions that manipulate whole tables, including crosstab';


--
-- Name: unaccent; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS unaccent WITH SCHEMA public;


--
-- Name: EXTENSION unaccent; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION unaccent IS 'text search dictionary that removes accents';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: notify_messenger_messages(); Type: FUNCTION; Schema: public; Owner: arthur
--

CREATE FUNCTION public.notify_messenger_messages() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    BEGIN
        PERFORM pg_notify('messenger_messages', NEW.queue_name::text);
        RETURN NEW;
    END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: character; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."character" (
    id integer NOT NULL,
    name text NOT NULL,
    gender text,
    ability text NOT NULL,
    minimal_distance numeric NOT NULL,
    weight numeric,
    born timestamp without time zone NOT NULL,
    in_space_since timestamp without time zone,
    beer_consumption integer NOT NULL,
    knows_the_answer boolean NOT NULL
);


ALTER TABLE public."character" OWNER TO postgres;

--
-- Name: character_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.character_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.character_id_seq OWNER TO postgres;

--
-- Name: character_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.character_id_seq OWNED BY public."character".id;


--
-- Name: messenger_messages; Type: TABLE; Schema: public; Owner: arthur
--

CREATE TABLE public.messenger_messages (
    id bigint NOT NULL,
    body text NOT NULL,
    headers text NOT NULL,
    queue_name character varying(190) NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    available_at timestamp(0) without time zone NOT NULL,
    delivered_at timestamp(0) without time zone DEFAULT NULL::timestamp without time zone
);


--
-- Name: messenger_messages_id_seq; Type: SEQUENCE; Schema: public; Owner: arthur
--

CREATE SEQUENCE public.messenger_messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: messenger_messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: arthur
--

ALTER SEQUENCE public.messenger_messages_id_seq OWNED BY public.messenger_messages.id;


--
-- Name: nemesis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.nemesis (
    is_alive boolean NOT NULL,
    years integer,
    id integer NOT NULL,
    character_id integer
);


ALTER TABLE public.nemesis OWNER TO postgres;

--
-- Name: nemesis_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.nemesis_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.nemesis_id_seq OWNER TO postgres;

--
-- Name: nemesis_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.nemesis_id_seq OWNED BY public.nemesis.id;


--
-- Name: secret; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.secret (
    id integer NOT NULL,
    secret_code bigint NOT NULL,
    nemesis_id integer NOT NULL
);


ALTER TABLE public.secret OWNER TO postgres;

--
-- Name: secrete_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.secrete_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.secrete_id_seq OWNER TO postgres;

--
-- Name: secrete_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.secrete_id_seq OWNED BY public.secret.id;


--
-- Name: character id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."character" ALTER COLUMN id SET DEFAULT nextval('public.character_id_seq'::regclass);


--
-- Name: messenger_messages id; Type: DEFAULT; Schema: public; Owner: arthur
--

ALTER TABLE ONLY public.messenger_messages ALTER COLUMN id SET DEFAULT nextval('public.messenger_messages_id_seq'::regclass);


--
-- Name: nemesis id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.nemesis ALTER COLUMN id SET DEFAULT nextval('public.nemesis_id_seq'::regclass);


--
-- Name: secret id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.secret ALTER COLUMN id SET DEFAULT nextval('public.secrete_id_seq'::regclass);


--
-- Data for Name: character; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."character" (id, name, gender, ability, minimal_distance, weight, born, in_space_since, beer_consumption, knows_the_answer) FROM stdin;
2	Trillian	female	mathematician	6.2	49	1994-12-14 00:00:00	2014-12-24 17:21:50	6704	t
4	Zaphod Beeblebrox	m	semi_half_cousin	1.6	91	1985-02-17 00:00:00	2014-12-04 04:09:55	679420	t
5	Zaphod Beeblebrox	m	semi_half_cousin	1.6	91	1985-02-17 00:00:00	2014-12-04 04:09:55	679242	t
6	Ford Prefect	M	has_towel	0.8	107	1990-05-04 00:00:00	2014-12-21 11:57:06	62544	t
10	Slartibartfast	male	coastlines_creator	0.4	96	1997-05-04 00:00:00	2014-09-19 01:26:41	46901	t
11	Arthur Dent	M	enjoys_tea	1.5	86	1989-09-02 00:00:00	2014-11-12 01:22:57	579	t
12	Frankie	mouse	NULL	4.0	0.18	1995-05-29 00:00:00	2014-11-27 12:46:16	55	t
13	Benjy	mouse	NULL	2.5	0.12	1989-04-02 00:00:00	2014-11-24 09:50:45	56	t
16	Eddie	M	quite_talkative	2.2	100	1989-04-04 00:00:00	2014-12-25 04:44:17	0	f
15	Deep Thought	\N	42	4.2	420	2000-02-04 00:42:00	2000-10-24 20:40:00	4242	t
9	Alice Beeblebrox	F	mothering	4.0	\N	1901-05-08 00:00:00	\N	64	f
\.


--
-- Data for Name: messenger_messages; Type: TABLE DATA; Schema: public; Owner: arthur
--

COPY public.messenger_messages (id, body, headers, queue_name, created_at, available_at, delivered_at) FROM stdin;
\.


--
-- Data for Name: nemesis; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.nemesis (is_alive, years, id, character_id) FROM stdin;
t	29	1	2
t	28	2	6
t	45	3	9
t	44	4	10
f	44	5	10
t	33	6	11
t	953	7	12
t	2	8	12
t	46	10	13
t	49	13	15
f	48	14	15
t	44	15	16
t	\N	11	13
\.


--
-- Data for Name: secret; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.secret (id, secret_code, nemesis_id) FROM stdin;
2	4168664804	1
3	5464826016	2
4	6270976449	2
5	7899028241	2
6	9442445871	2
7	9449428626	3
8	4169606040	3
9	4167477856	3
10	4168664804	4
11	5464646769	5
12	8422644094	6
13	5467717091	7
14	4166492176	7
15	6271440484	8
16	6275689247	8
19	6275900044	10
20	8422745472	10
21	8467224	10
22	6277097442	11
23	6272449905	11
24	8424742058	14
25	1798274556	14
26	415701	15
27	7898402279	15
28	1246020766	15
\.


--
-- Name: character character_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."character"
    ADD CONSTRAINT character_pkey PRIMARY KEY (id);


--
-- Name: messenger_messages messenger_messages_pkey; Type: CONSTRAINT; Schema: public; Owner: arthur
--

ALTER TABLE ONLY public.messenger_messages
    ADD CONSTRAINT messenger_messages_pkey PRIMARY KEY (id);


--
-- Name: nemesis nemesis_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.nemesis
    ADD CONSTRAINT nemesis_pkey PRIMARY KEY (id);


--
-- Name: secret secrete_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.secret
    ADD CONSTRAINT secrete_pkey PRIMARY KEY (id);


--
-- Name: fki_Character Id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "fki_Character Id" ON public.nemesis USING btree (id);


--
-- Name: idx_75ea56e016ba31db; Type: INDEX; Schema: public; Owner: arthur
--

CREATE INDEX idx_75ea56e016ba31db ON public.messenger_messages USING btree (delivered_at);


--
-- Name: idx_75ea56e0e3bd61ce; Type: INDEX; Schema: public; Owner: arthur
--

CREATE INDEX idx_75ea56e0e3bd61ce ON public.messenger_messages USING btree (available_at);


--
-- Name: idx_75ea56e0fb7336f0; Type: INDEX; Schema: public; Owner: arthur
--

CREATE INDEX idx_75ea56e0fb7336f0 ON public.messenger_messages USING btree (queue_name);


--
-- Name: messenger_messages notify_trigger; Type: TRIGGER; Schema: public; Owner: arthur
--

CREATE TRIGGER notify_trigger AFTER INSERT OR UPDATE ON public.messenger_messages FOR EACH ROW EXECUTE FUNCTION public.notify_messenger_messages();


--
-- Name: nemesis character; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.nemesis
    ADD CONSTRAINT "character" FOREIGN KEY (character_id) REFERENCES public."character"(id) NOT VALID;


--
-- Name: secret nemesis; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.secret
    ADD CONSTRAINT nemesis FOREIGN KEY (nemesis_id) REFERENCES public.nemesis(id);

--
-- PostgreSQL database dump complete
--

