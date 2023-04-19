// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB cluster from a DB snapshot or DB cluster snapshot. The target
// DB cluster is created from the source snapshot with a default configuration. If
// you don't specify a security group, the new DB cluster is associated with the
// default security group. This action only restores the DB cluster, not the DB
// instances for that DB cluster. You must invoke the CreateDBInstance action to
// create DB instances for the restored DB cluster, specifying the identifier of
// the restored DB cluster in DBClusterIdentifier . You can create DB instances
// only after the RestoreDBClusterFromSnapshot action has completed and the DB
// cluster is available. For more information on Amazon Aurora DB clusters, see
// What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) RestoreDBClusterFromSnapshot(ctx context.Context, params *RestoreDBClusterFromSnapshotInput, optFns ...func(*Options)) (*RestoreDBClusterFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreDBClusterFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterFromSnapshot", params, optFns, c.addOperationRestoreDBClusterFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBClusterFromSnapshotInput struct {

	// The name of the DB cluster to create from the DB snapshot or DB cluster
	// snapshot. This parameter isn't case-sensitive. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens
	//   - First character must be a letter
	//   - Can't end with a hyphen or contain two consecutive hyphens
	// Example: my-snapshot-id Valid for: Aurora DB clusters and Multi-AZ DB clusters
	//
	// This member is required.
	DBClusterIdentifier *string

	// The database engine to use for the new DB cluster. Default: The same as source
	// Constraint: Must be compatible with the engine of the source Valid for: Aurora
	// DB clusters and Multi-AZ DB clusters
	//
	// This member is required.
	Engine *string

	// The identifier for the DB snapshot or DB cluster snapshot to restore from. You
	// can use either the name or the Amazon Resource Name (ARN) to specify a DB
	// cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
	// Constraints:
	//   - Must match the identifier of an existing Snapshot.
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	//
	// This member is required.
	SnapshotIdentifier *string

	// Provides the list of Availability Zones (AZs) where instances in the restored
	// DB cluster can be created. Valid for: Aurora DB clusters only
	AvailabilityZones []string

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0. Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	// Default: 0 Constraints:
	//   - If specified, this value must be set to a number from 0 to 259,200 (72
	//   hours).
	// Valid for: Aurora DB clusters only
	BacktrackWindow *int64

	// A value that indicates whether to copy all tags from the restored DB cluster to
	// snapshots of the restored DB cluster. The default is not to copy them. Valid
	// for: Aurora DB clusters and Multi-AZ DB clusters
	CopyTagsToSnapshot *bool

	// The compute and memory capacity of the each DB instance in the Multi-AZ DB
	// cluster, for example db.m6gd.xlarge. Not all DB instance classes are available
	// in all Amazon Web Services Regions, or for all database engines. For the full
	// list of DB instance classes, and availability for your engine, see DB Instance
	// Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Valid for: Multi-AZ DB clusters only
	DBClusterInstanceClass *string

	// The name of the DB cluster parameter group to associate with this DB cluster.
	// If this argument is omitted, the default DB cluster parameter group for the
	// specified engine is used. Constraints:
	//   - If supplied, must match the name of an existing default DB cluster
	//   parameter group.
	//   - Must be 1 to 255 letters, numbers, or hyphens.
	//   - First character must be a letter.
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	DBClusterParameterGroupName *string

	// The name of the DB subnet group to use for the new DB cluster. Constraints: If
	// supplied, must match the name of an existing DB subnet group. Example:
	// mydbsubnetgroup Valid for: Aurora DB clusters and Multi-AZ DB clusters
	DBSubnetGroupName *string

	// The database name for the restored DB cluster. Valid for: Aurora DB clusters
	// and Multi-AZ DB clusters
	DatabaseName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	DeletionProtection *bool

	// Specify the Active Directory directory ID to restore the DB cluster in. The
	// domain must be created prior to this operation. Currently, only MySQL, Microsoft
	// SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active
	// Directory Domain. For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)
	// in the Amazon RDS User Guide. Valid for: Aurora DB clusters only
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service. Valid for: Aurora DB clusters only
	DomainIAMRoleName *string

	// The list of logs that the restored DB cluster is to export to Amazon CloudWatch
	// Logs. The values in the list depend on the DB engine being used. RDS for MySQL
	// Possible values are error , general , and slowquery . RDS for PostgreSQL
	// Possible values are postgresql and upgrade . Aurora MySQL Possible values are
	// audit , error , general , and slowquery . Aurora PostgreSQL Possible value is
	// postgresql . For more information about exporting CloudWatch Logs for Amazon
	// RDS, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. For more information about exporting CloudWatch
	// Logs for Amazon Aurora, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	EnableCloudwatchLogsExports []string

	// A value that indicates whether to enable mapping of Amazon Web Services
	// Identity and Access Management (IAM) accounts to database accounts. By default,
	// mapping isn't enabled. For more information, see IAM Database Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	EnableIAMDatabaseAuthentication *bool

	// The DB engine mode of the DB cluster, either provisioned or serverless . For
	// more information, see CreateDBCluster (https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBCluster.html)
	// . Valid for: Aurora DB clusters only
	EngineMode *string

	// The version of the database engine to use for the new DB cluster. If you don't
	// specify an engine version, the default version for the database engine in the
	// Amazon Web Services Region is used. To list all of the available engine versions
	// for Aurora MySQL, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for Aurora PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-postgresql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for MySQL, use the following command: aws rds
	// describe-db-engine-versions --engine mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine postgres --query
	// "DBEngineVersions[].EngineVersion" Aurora MySQL See Database engine updates for
	// Amazon Aurora MySQL (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html)
	// in the Amazon Aurora User Guide. Aurora PostgreSQL See Amazon Aurora PostgreSQL
	// releases and engine versions (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html)
	// in the Amazon Aurora User Guide. MySQL See Amazon RDS for MySQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt)
	// in the Amazon RDS User Guide. PostgreSQL See Amazon RDS for PostgreSQL versions
	// and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts)
	// in the Amazon RDS User Guide. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	EngineVersion *string

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for each DB instance in the Multi-AZ DB cluster. For
	// information about valid IOPS values, see Amazon RDS Provisioned IOPS storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide. Constraints: Must be a multiple between .5 and 50
	// of the storage amount for the DB instance. Valid for: Aurora DB clusters and
	// Multi-AZ DB clusters
	Iops *int32

	// The Amazon Web Services KMS key identifier to use when restoring an encrypted
	// DB cluster from a DB snapshot or DB cluster snapshot. The Amazon Web Services
	// KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS
	// key. To use a KMS key in a different Amazon Web Services account, specify the
	// key ARN or alias ARN. When you don't specify a value for the KmsKeyId
	// parameter, then the following occurs:
	//   - If the DB snapshot or DB cluster snapshot in SnapshotIdentifier is
	//   encrypted, then the restored DB cluster is encrypted using the KMS key that was
	//   used to encrypt the DB snapshot or DB cluster snapshot.
	//   - If the DB snapshot or DB cluster snapshot in SnapshotIdentifier isn't
	//   encrypted, then the restored DB cluster isn't encrypted.
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	KmsKeyId *string

	// The network type of the DB cluster. Valid values:
	//   - IPV4
	//   - DUAL
	// The network type is determined by the DBSubnetGroup specified for the DB
	// cluster. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the
	// IPv6 protocols ( DUAL ). For more information, see  Working with a DB instance
	// in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	NetworkType *string

	// The name of the option group to use for the restored DB cluster. DB clusters
	// are associated with a default option group that can't be modified.
	OptionGroupName *string

	// The port number on which the new DB cluster accepts connections. Constraints:
	// This value must be 1150-65535 Default: The same port as the original DB
	// cluster. Valid for: Aurora DB clusters and Multi-AZ DB clusters
	Port *int32

	// A value that indicates whether the DB cluster is publicly accessible. When the
	// DB cluster is publicly accessible, its Domain Name System (DNS) endpoint
	// resolves to the private IP address from within the DB cluster's virtual private
	// cloud (VPC). It resolves to the public IP address from outside of the DB
	// cluster's VPC. Access to the DB cluster is ultimately controlled by the security
	// group it uses. That public access is not permitted if the security group
	// assigned to the DB cluster doesn't permit it. When the DB cluster isn't publicly
	// accessible, it is an internal DB cluster with a DNS name that resolves to a
	// private IP address. Default: The default behavior varies depending on whether
	// DBSubnetGroupName is specified. If DBSubnetGroupName isn't specified, and
	// PubliclyAccessible isn't specified, the following applies:
	//   - If the default VPC in the target Region doesn’t have an internet gateway
	//   attached to it, the DB cluster is private.
	//   - If the default VPC in the target Region has an internet gateway attached to
	//   it, the DB cluster is public.
	// If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the
	// following applies:
	//   - If the subnets are part of a VPC that doesn’t have an internet gateway
	//   attached to it, the DB cluster is private.
	//   - If the subnets are part of a VPC that has an internet gateway attached to
	//   it, the DB cluster is public.
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	PubliclyAccessible *bool

	// For DB clusters in serverless DB engine mode, the scaling properties of the DB
	// cluster. Valid for: Aurora DB clusters only
	ScalingConfiguration *types.ScalingConfiguration

	// Contains the scaling configuration of an Aurora Serverless v2 DB cluster. For
	// more information, see Using Amazon Aurora Serverless v2 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html)
	// in the Amazon Aurora User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// Specifies the storage type to be associated with the each DB instance in the
	// Multi-AZ DB cluster. Valid values: io1 When specified, a value for the Iops
	// parameter is required. Default: io1 Valid for: Aurora DB clusters and Multi-AZ
	// DB clusters
	StorageType *string

	// The tags to be assigned to the restored DB cluster. Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	Tags []types.Tag

	// A list of VPC security groups that the new DB cluster will belong to. Valid
	// for: Aurora DB clusters and Multi-AZ DB clusters
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBClusterFromSnapshotOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , PromoteReadReplicaDBCluster ,
	// RestoreDBClusterFromS3 , RestoreDBClusterFromSnapshot ,
	// RestoreDBClusterToPointInTime , StartDBCluster , and StopDBCluster . For a
	// Multi-AZ DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , RebootDBCluster ,
	// RestoreDBClusterFromSnapshot , and RestoreDBClusterToPointInTime . For more
	// information on Amazon Aurora DB clusters, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see Multi-AZ deployments with two readable standby DB instances (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBClusterFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRestoreDBClusterFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRestoreDBClusterFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBClusterFromSnapshot",
	}
}
