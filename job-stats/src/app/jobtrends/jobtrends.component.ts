import { Component, OnInit, ViewChild } from '@angular/core';
import { PrimeNGConfig } from 'primeng/api';
import { Table } from 'primeng/table';
import { JobStats } from 'src/services/jobstats/jobstats';
import { JobstatsService } from 'src/services/jobstats/jobstats.service';

@Component({
    selector: 'app-jobtrends',
    templateUrl: './jobtrends.component.html',
    styleUrls: ['./jobtrends.component.scss']
})


export class JobtrendsComponent implements OnInit {
    basicData: any;
    basicOptions: {
        plugins: {
            legend: {
                labels: {
                    color: '#495057'
                }
            }
        },
        scales: {
            x: {
                ticks: {
                    color: '#495057'
                },
                grid: {
                    color: '#ebedef'
                }
            },
            y: {
                ticks: {
                    color: '#495057'
                },
                grid: {
                    color: '#ebedef'
                }
            }
        }
    };
    jobList: any[] = [];
    selectedJobDetails: JobStats[] = [];
    loading: boolean = true;
    selectedJobName: any;
    selectedJob: any;

    @ViewChild('dt') table: Table;

    constructor(private jobStatsService: JobstatsService, private primengConfig: PrimeNGConfig) {
    }

    ngOnInit(): void {
        this.jobStatsService.getUniqueJobNames().then(jobs => {
            this.jobList = jobs;
            this.loading = false;
        });
    }

    jobSelectionChanged(event: any) {
        this.selectedJobName = event.value;
        this.refreshChartData()

    }
    refreshChartData() {
        // for (const selectedJob of this.selectedJobNames) {
        //     this.jobStatsService.getJobDetails(selectedJob).then(jobs => {
        //         this.selectedJobInfo.push(jobs);

        //     });
        // }
        this.jobStatsService.getJobDetails(this.selectedJobName).then(jobDetail => {
            this.selectedJobDetails = jobDetail;
            const runDates: string[] = this.selectedJobDetails.map(stats => stats.runDate);
            const runTimes: string[] = this.selectedJobDetails.map(stats => stats.runTime);
            this.basicData = {
                labels: runDates,
                datasets: [
                    {
                        label: this.selectedJobName,
                        data: runTimes,
                        fill: false,
                        borderColor: '#42A5F5',
                        tension: .4
                    }
                ]
            };
        });
    }
}
