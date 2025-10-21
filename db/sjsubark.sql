--
-- PostgreSQL database dump
--

\restrict grHROJtfCBGjPzFhUJ9mLQjqWx4nOauKG1crnroWEalscIsAmM7Sb46PqYnAKrK

-- Dumped from database version 18.0
-- Dumped by pg_dump version 18.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: garage; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.garage (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    addr character varying(255) NOT NULL
);


ALTER TABLE public.garage OWNER TO postgres;

--
-- Name: garage_fullness; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.garage_fullness (
    transaction_id integer NOT NULL,
    name character varying(50) NOT NULL,
    utc_timestamp timestamp without time zone NOT NULL,
    second integer NOT NULL,
    minute integer NOT NULL,
    hour integer NOT NULL,
    day integer NOT NULL,
    month integer NOT NULL,
    year integer NOT NULL,
    weekday integer NOT NULL,
    is_weekend boolean NOT NULL,
    is_campus_closed boolean NOT NULL
);


ALTER TABLE public.garage_fullness OWNER TO postgres;

--
-- Name: garage_fullness_transaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.garage_fullness_transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.garage_fullness_transaction_id_seq OWNER TO postgres;

--
-- Name: garage_fullness_transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.garage_fullness_transaction_id_seq OWNED BY public.garage_fullness.transaction_id;


--
-- Name: garage_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.garage_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.garage_id_seq OWNER TO postgres;

--
-- Name: garage_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.garage_id_seq OWNED BY public.garage.id;


--
-- Name: garage id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.garage ALTER COLUMN id SET DEFAULT nextval('public.garage_id_seq'::regclass);


--
-- Name: garage_fullness transaction_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.garage_fullness ALTER COLUMN transaction_id SET DEFAULT nextval('public.garage_fullness_transaction_id_seq'::regclass);


--
-- Name: garage_fullness garage_fullness_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.garage_fullness
    ADD CONSTRAINT garage_fullness_pkey PRIMARY KEY (transaction_id);


--
-- Name: garage garage_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.garage
    ADD CONSTRAINT garage_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

\unrestrict grHROJtfCBGjPzFhUJ9mLQjqWx4nOauKG1crnroWEalscIsAmM7Sb46PqYnAKrK

