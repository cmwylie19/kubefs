```shell
jobs:
  Job1:
    runs-on: ubuntu-latest
    steps:
    - name: Do something
      run: echo "Job 1 running"

  Job2:
    runs-on: ubuntu-latest
    steps:
    - name: Do something
      run: echo "Job 2 running"

  Job3:
    runs-on: ubuntu-latest
    steps:
    - name: Do something
      run: echo "Job 3 running"
 
  Job4:
    if: success()
    needs: [Job1, Job2, Job3]
    runs-on: ubuntu-latest
    steps:
    - name: Do something
      run: echo "Job 4 running"
```