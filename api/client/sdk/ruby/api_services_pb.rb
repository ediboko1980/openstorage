# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: api.proto for package 'openstorage.api'

require 'grpc'
require 'api_pb'

module Openstorage
  module Api
    module OpenStorageCluster
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'openstorage.api.OpenStorageCluster'

        # Enumerate lists all the nodes in the cluster.
        rpc :Enumerate, SdkClusterEnumerateRequest, SdkClusterEnumerateResponse
        # Inspect the node given a UUID.
        rpc :Inspect, SdkClusterInspectRequest, SdkClusterInspectResponse
        # Get a list of alerts from the storage cluster
        rpc :AlertEnumerate, SdkClusterAlertEnumerateRequest, SdkClusterAlertEnumerateResponse
        # Clear the alert for a given resource
        rpc :AlertClear, SdkClusterAlertClearRequest, SdkClusterAlertClearResponse
        # Erases an alert for a given resource
        rpc :AlertDelete, SdkClusterAlertDeleteRequest, SdkClusterAlertDeleteResponse
      end

      Stub = Service.rpc_stub_class
    end
    module OpenStorageVolume
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'openstorage.api.OpenStorageVolume'

        # Creates a new volume
        rpc :Create, SdkVolumeCreateRequest, SdkVolumeCreateResponse
        # Clone creates a new volume cloned from an existing volume
        rpc :Clone, SdkVolumeCloneRequest, SdkVolumeCloneResponse
        # Delete a volume
        rpc :Delete, SdkVolumeDeleteRequest, SdkVolumeDeleteResponse
        # Get information on a volume
        rpc :Inspect, SdkVolumeInspectRequest, SdkVolumeInspectResponse
        # Get a list of volumes
        rpc :Enumerate, SdkVolumeEnumerateRequest, SdkVolumeEnumerateResponse
        # Create a snapshot of a volume. This creates an immutable (read-only),
        # point-in-time snapshot of a volume.
        rpc :SnapshotCreate, SdkVolumeSnapshotCreateRequest, SdkVolumeSnapshotCreateResponse
        # Restores a volume to a specified snapshot
        rpc :SnapshotRestore, SdkVolumeSnapshotRestoreRequest, SdkVolumeSnapshotRestoreResponse
        # List the number of snapshots for a specific volume
        rpc :SnapshotEnumerate, SdkVolumeSnapshotEnumerateRequest, SdkVolumeSnapshotEnumerateResponse
        # Attach device to host
        rpc :Attach, SdkVolumeAttachRequest, SdkVolumeAttachResponse
        # Detaches the volume from the node.
        rpc :Detach, SdkVolumeDetachRequest, SdkVolumeDetachResponse
        # Attaches the volume to a node.
        rpc :Mount, SdkVolumeMountRequest, SdkVolumeMountResponse
        # Unmount volume at specified path
        rpc :Unmount, SdkVolumeUnmountRequest, SdkVolumeUnmountResponse
      end

      Stub = Service.rpc_stub_class
    end
    module OpenStorageObjectstore
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'openstorage.api.OpenStorageObjectstore'

        # Inspect returns current status of objectstore
        rpc :Inspect, SdkObjectstoreInspectRequest, SdkObjectstoreInspectResponse
        # Creates objectstore on specified volume
        rpc :Create, SdkObjectstoreCreateRequest, SdkObjectstoreCreateResponse
        # Deletes objectstore by id
        rpc :Delete, SdkObjectstoreDeleteRequest, SdkObjectstoreDeleteResponse
        # Updates provided objectstore status
        rpc :Update, SdkObjectstoreUpdateRequest, SdkObjectstoreUpdateResponse
      end

      Stub = Service.rpc_stub_class
    end
    module OpenStorageCredentials
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'openstorage.api.OpenStorageCredentials'

        # Provide credentials to OpenStorage and if valid,
        # it will return an identifier to the credentials
        #
        # Create cloud credentials
        rpc :Create, SdkCredentialCreateRequest, SdkCredentialCreateResponse
        # Enumerate returns a list of credential ids
        rpc :Enumerate, SdkCredentialEnumerateRequest, SdkCredentialEnumerateResponse
        # Inspect returns the information about a credential
        rpc :Inspect, SdkCredentialInspectRequest, SdkCredentialInspectResponse
        # Delete a specified credential
        rpc :Delete, SdkCredentialDeleteRequest, SdkCredentialDeleteResponse
        # Validate a specified credential
        rpc :Validate, SdkCredentialValidateRequest, SdkCredentialValidateResponse
      end

      Stub = Service.rpc_stub_class
    end
    module OpenStorageSchedulePolicy
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'openstorage.api.OpenStorageSchedulePolicy'

        # Create Schedule Policy for snapshots
        rpc :Create, SdkSchedulePolicyCreateRequest, SdkSchedulePolicyCreateResponse
        # Update Schedule Policy
        rpc :Update, SdkSchedulePolicyUpdateRequest, SdkSchedulePolicyUpdateResponse
        rpc :Enumerate, SdkSchedulePolicyEnumerateRequest, SdkSchedulePolicyEnumerateResponse
        # Inspect Schedule Policy
        rpc :Inspect, SdkSchedulePolicyInspectRequest, SdkSchedulePolicyInspectResponse
        # Delete Schedule Policy
        rpc :Delete, SdkSchedulePolicyDeleteRequest, SdkSchedulePolicyDeleteResponse
      end

      Stub = Service.rpc_stub_class
    end
    module OpenStorageCloudBackup
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'openstorage.api.OpenStorageCloudBackup'

        # Create
        rpc :Create, SdkCloudBackupCreateRequest, SdkCloudBackupCreateResponse
        # Restore
        rpc :Restore, SdkCloudBackupRestoreRequest, SdkCloudBackupRestoreResponse
        # Delete
        rpc :Delete, SdkCloudBackupDeleteRequest, SdkCloudBackupDeleteResponse
        # DeleteAll
        rpc :DeleteAll, SdkCloudBackupDeleteAllRequest, SdkCloudBackupDeleteAllResponse
        # Enumerate
        rpc :Enumerate, SdkCloudBackupEnumerateRequest, SdkCloudBackupEnumerateResponse
        # Status
        rpc :Status, SdkCloudBackupStatusRequest, SdkCloudBackupStatusResponse
        # Catalog
        rpc :Catalog, SdkCloudBackupCatalogRequest, SdkCloudBackupCatalogResponse
        # History
        rpc :History, SdkCloudBackupHistoryRequest, SdkCloudBackupHistoryResponse
        # StateChange
        rpc :StateChange, SdkCloudBackupStateChangeRequest, SdkCloudBackupStateChangeResponse
        # Create Schedule
        rpc :SchedCreate, SdkCloudBackupSchedCreateRequest, SdkCloudBackupSchedCreateResponse
        # Delete Schedule
        rpc :SchedDelete, SdkCloudBackupSchedDeleteRequest, SdkCloudBackupSchedDeleteResponse
        # Enumerate schedules
        rpc :SchedEnumerate, SdkCloudBackupSchedEnumerateRequest, SdkCloudBackupSchedEnumerateResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
