package v1

import (
	"net/url"
	"strings"

	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/kubescape/v2/core/cautils/getter"
	"github.com/armosec/opa-utils/reporthandling"
	"github.com/google/uuid"
)

func (report *ReportEventReceiver) initEventReceiverURL() {
	urlObj := parseHost(getter.GetArmoAPIConnector().GetReportReceiverURL())

	urlObj.Path = "/k8s/postureReport"
	q := urlObj.Query()
	q.Add("customerGUID", uuid.MustParse(report.customerGUID).String())
	q.Add("clusterName", report.clusterName)

	urlObj.RawQuery = q.Encode()

	report.eventReceiverURL = &urlObj
}
func parseHost(host string) url.URL {
	urlObj := url.URL{}
	if strings.Contains(host, "http://") {
		urlObj.Scheme = "http"
		urlObj.Host = strings.Replace(host, "http://", "", 1)
	} else {
		urlObj.Scheme = "https"
		urlObj.Host = host
	}
	return urlObj
}
func hostToString(host *url.URL, reportID string) string {
	q := host.Query()
	q.Add("reportID", reportID) // TODO - do we add the reportID?
	host.RawQuery = q.Encode()
	return host.String()
}

func setPaginationReport(postureReport *reporthandling.PostureReport) *reporthandling.PostureReport {
	return &reporthandling.PostureReport{
		CustomerGUID:         postureReport.CustomerGUID,
		ClusterName:          postureReport.ClusterName,
		ReportID:             postureReport.ReportID,
		ReportGenerationTime: postureReport.ReportGenerationTime,
	}
}
func iMetaToResource(obj workloadinterface.IMetadata) *reporthandling.Resource {
	return &reporthandling.Resource{
		ResourceID: obj.GetID(),
		Object:     obj.GetObject(),
	}
}
