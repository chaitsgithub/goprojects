import { Component, OnInit, ViewChild } from '@angular/core';
import {JobStats} from '../../services/jobstats/jobstats'
import {JobstatsService} from '../../services/jobstats/jobstats.service'
import { Table } from 'primeng/table';
import { PrimeNGConfig } from 'primeng/api';

@Component({
  selector: 'app-jobstats',
  templateUrl: './jobstats.component.html',
  styleUrls: ['./jobstats.component.scss']
})
export class JobstatsComponent {
  [x: string]: any; 
  jobstats: JobStats[] = [];
  selectedJobs: JobStats[] = [];
  statuses: any[] = [];
  loading: boolean = true;

  @ViewChild('dt') table: Table;

  constructor(private jobStatsService: JobstatsService, private primengConfig: PrimeNGConfig) {
  }

  ngOnInit() {
      this.jobStatsService.getJobStatsLarge().then(jobstats => {
          this.jobstats = jobstats;
          this.loading = false;
      });

      this.jobs = ["Job#1","Job#2","Job#3","Job#4"];

      this.statuses = [
          {label: 'running', value: 'warning'},
          {label: 'success', value: 'success'},
          {label: 'fail', value: 'danger'}
      ]
      this.primengConfig.ripple = true;
  }

  onDateSelect(value) {
      this.table.filter(this.formatDate(value), 'date', 'equals')
  }
}
