import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JobStats } from './jobstats';

@Injectable()
export class JobstatsService {

  constructor(private http: HttpClient) { }

  getJobStatsLarge() {
    return this.http.get<any>('assets/jobstats-large.json')
      .toPromise()
      .then(res => <JobStats[]>res.data)
      .then(data => { return data; });
  }

  getUniqueJobNames() {
    return this.http.get<any>('assets/jobstats-large.json')
      .toPromise()
      .then(res => <JobStats[]>res.data)
      .then(data => {
        const jobNames = new Set<string>();
        data.forEach(job => {
          if (job.jobName) {
            jobNames.add(job.jobName);
          }
        });
        return Array.from(jobNames);
      });
  }
  getJobDetails(selectedJobName) {
    return this.http.get<any>('assets/jobstats-large.json')
      .toPromise()
      .then(res => <JobStats[]>res.data)
      .then(data => data.filter(job => job.jobName === selectedJobName));
  }
}
