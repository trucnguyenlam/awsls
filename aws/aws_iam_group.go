// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamGroup(client *Client) ([]Resource, error) {
	req := client.iamconn.ListGroupsRequest(&iam.ListGroupsInput{})

	var result []Resource

	p := iam.NewListGroupsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Groups {

			t := *r.CreateDate
			result = append(result, Resource{
				Type:   "aws_iam_group",
				ID:     *r.GroupName,
				Region: client.iamconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
