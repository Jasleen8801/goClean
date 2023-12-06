import { useState, useEffect } from "react";
import axios from "axios";

const url = "http://localhost:3000";

const useFetch = () => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchData = async (endpoint, method, requestData = null) => {
    setLoading(true);
    
    try {
      let response;

      if (method === "GET" || method === "DELETE") {
        response = await axios[method.toLowerCase()](`${url}/${endpoint}`);
      } else if (method === "POST" || method === "PUT") {
        response = await axios[method.toLowerCase()](`${url}/${endpoint}`, requestData);
      }

      setData(response.data);
      setLoading(false);
    } catch (error) {
      console.error(error);
      setError(error);
      setLoading(false);
    }
  };

  const get = async (endpoint) => {
    await fetchData(endpoint, "GET");
  };

  const post = async (endpoint, requestData) => {
    await fetchData(endpoint, "POST", requestData);
  };

  const put = async (endpoint, requestData) => {
    await fetchData(endpoint, "PUT", requestData);
  };

  const del = async (endpoint) => {
    await fetchData(endpoint, "DELETE");
  };

  useEffect(() => {
    get(); 
  }, []);

  return { data, loading, error, get, post, put, del };
};

export default useFetch;
