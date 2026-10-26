## Summary:
Rambler is a web application that allows users input (speak) to their computer or phone's microphone, and recieve a formatted and structured output. Is designed to be a Notes/Todo hybrid app which classifies user input as either a note or a todo. 


## MVP 
- Users can speak to their microphone and recieve structured outputs
- Tasks should be classified as general notes vs todos
- Todo's should be further classified based on the urgency and timeframe of the task (Today, This Week, This Month, This Year)
- Native Calender feature
- Speak tab, Notes Tab, Calender Tab
- Daily Task Queue

- Vector Database for semantic (Later?)
     - Ex someone cant exactly remember the keywords to search by but wants to search by an idea, "when i was talking about that math problem "


## How it Works
- A user speaks into their microphone.
- Their speecg is converted to text by an open source whisper model.
- text blob is fed to the server.
- the server sends the text blob off to first run through embedding model (for later) than an open srouce LLM being hosted on a remote machine using vllm.
- The LLM analyzed the text blob and formats it based on a prompt I've given it. 
- The LLM returns the formatted outputs to the server.
- The server forwards the outputs back to the frontend where they are displayed.


## Objectives
- Solidify understanding of writing a Go Backend
- Solidify understanding of serving local models via VLLM
- Solidfy understanding of agentic workflows on the frontend by having agents design an intuitive interface for both web and mobile.

## Stack
 - Vue or Angular?
 - Browser Native MediaStream API for speech recording.
 - Go
 - VLLM
 - Docker
 - Qwen2.5-14B / Qwen3-14B - Qwen/Qwen2.5-14B-Instruct (or Qwen3-14B-Instruct if available) - AWQ 4-bit or GGUF Q4/Q5
 - JSON


 ## To USE
 - To access from a laptop
    1. Serve the model on the remote machine as per usual : vllm serve Qwen/Qwen2.5-14B-Instruct-AWQ --max-model-len 8192
    2. from the laptop open an SSH tunnel : ssh -N -L 8000:localhost:8000 user@your-remote-ip, it will look like it hanging
    3. open a new terminal and test: curl http://localhost:8000/v1/models
    4. response should look something like: {"object":"list","data":[{"id":"Qwen/Qwen2.5-14B-Instruct-AWQ","object":"model","created":1792688980,"owned_by":"vllm","root":"Qwen/Qwen2.5-14B-Instruct-AWQ","parent":null,"max_model_len":8192,"permission":[{"id":"modelperm-82a27e7747a1a2b2","object":"model_permission","created":1792688980,"allow_create_engine":false,"allow_sampling":true,"allow_logprobs":true,"allow_search_indices":false,"allow_view":true,"allow_fine_tuning":false,"organization":"*","group":null,"is_blocking":false}]}]}% 

- Testing with an audio file use curl.
     curl -X POST "http://127.0.0.1:9000/transcribe" \
     -F "file=@tests/testrecord1.m4a"

